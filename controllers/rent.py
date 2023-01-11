from models import Rent, Status 

from controllers.rentType import (
    get_rentType_by_id,
    get_rentType_daily_period,
    )


from controllers.lockers import(
    get_locker_id,
    rent_locker,
    release_locker
)

from controllers.transactionLog import cal_transaction_amount

from datetime import datetime,timedelta
from database import db 
from sqlalchemy.exc import SQLAlchemyError

def period_elapsed(rentType_id, rent_date_from, rent_date_to):
    type = get_rentType_by_id(rentType_id)

    if not type:
        return -1

    if type.type.value == "Hourly":
        time = rent_date_to - rent_date_from   
        return (time.days*24) + (time.seconds/3600)
    elif type.type.value == "Daily":
        time = rent_date_to - rent_date_from   
        return  time.days
    elif type.type.value == "Weekly":
        time = rent_date_to - rent_date_from   
        return time.days/7
    elif type.type.value == "Monthly":
        time = (rent_date_to.year - rent_date_from.year)  * 12 + (rent_date_to.month - rent_date_from.month)
        if time == 0 or time == -1:
            time = 1
        return  time
    elif type.type.value == "Semester":
        return 1

def init_amount_owed(rentType_id, rent_date_from, rent_date_to):
    type = get_rentType_by_id(rentType_id)

    if not type:
        return -1

    price = float(type.price)

    return price * period_elapsed(rentType_id, rent_date_from, rent_date_to)
        
def recal_amount_owed(rentType_id,date_returned,rent_date_from,rent_date_to):
    if date_returned:
        if date_returned > rent_date_to:
            return init_amount_owed(rentType_id,rent_date_from,rent_date_to) + late_fees(rentType_id,date_returned,rent_date_from,rent_date_to)

        if date_returned < rent_date_to:
            return init_amount_owed(rentType_id,rent_date_from,date_returned)
        return init_amount_owed(rentType_id,rent_date_from,date_returned)
    elif not date_returned:
        timestamp = datetime.now()
        if timestamp > rent_date_to:
           return init_amount_owed(rentType_id,rent_date_from,rent_date_to) + late_fees(rentType_id,timestamp,rent_date_from,rent_date_to)
        return init_amount_owed(rentType_id,rent_date_from,rent_date_to)

def late_fees(rentType_id, date_returned, rent_date_from, rent_date_to):
    type = get_rentType_by_id(rentType_id)
    
    if not type:
        return -1
    daily_price = get_rentType_daily_period(rent_date_from,rent_date_to)
    
    if not daily_price:
        return -1
    timestamp = datetime.now()
    if type.type.value == "Hourly":
        if timestamp > rent_date_to and not date_returned:
            time = period_elapsed(rentType_id, rent_date_to, timestamp)
            
        elif date_returned and date_returned > rent_date_to:
            time  = period_elapsed(rentType_id, rent_date_to, date_returned)
        else:
            return 0.00
        return type.price * time
    else:

        if timestamp > rent_date_to and not date_returned:
            time  = period_elapsed(daily_price.id, rent_date_to, timestamp)

        elif date_returned and date_returned > rent_date_to:
                time  = period_elapsed(daily_price.id, rent_date_to, date_returned)
        else:
            return 0.0
        return daily_price.price * time

def create_rent(student_id, locker_id,rentType, rent_date_from, rent_date_to):
    if get_overdue_rent_by_student(student_id) or get_owed_rent_by_student(student_id):
        return []

    if get_locker_id(locker_id): 
        try:
            amount_owed = init_amount_owed(rentType, rent_date_from, rent_date_to)
            rent = Rent(student_id, locker_id, rentType,rent_date_from,rent_date_to,amount_owed)
            db.session.add(rent)
            db.session.commit()
            rent_locker(locker_id)
            return rent
        except SQLAlchemyError as e:
            print(e)
            db.session.rollback()   
            return None
        


def get_rent_by_id(id):
    rent = Rent.query.filter_by(id=id).first()

    if not rent:
        return None
   
    return rent

def update_rent(id):
    rent = get_rent_by_id(id)

    if not rent:
        return None
    amt = recal_amount_owed(rent.rent_type,rent.date_returned,rent.rent_date_from,rent.rent_date_to)
    if amt is not None:
        rent.amount_owed = amt - cal_transaction_amount(id)
    rent.status = rent.check_status()

    try:
        db.session.add(rent)
        db.session.commit()
        return rent

    except SQLAlchemyError:
        db.session.rollback()
        return None


def get_overdue_rent_by_student(s_id):
    rent = Rent.query.filter_by(student_id= s_id,status = Status.OVERDUE).first()
    if not rent :
        return None
    return rent

def get_owed_rent_by_student(s_id):
    rent = Rent.query.filter_by(student_id= s_id,status = Status.OWED).first()
    if not rent :
        return None
    return rent

def release_rental(id,d_returned):
    rent = get_rent_by_id(id)

    if not rent:
        return None
    rent.date_returned = d_returned
    rent.status = Status.RETURNED
    
    try:
        db.session.add(rent)
        db.session.commit()
        release_locker(rent.locker_id)
        return rent

    except SQLAlchemyError:
        db.session.rollback()
        return None

def get_all_rentals():
    rents = Rent.query.all()

    if not rents:
        return None

    for r in rents:
        update_rent(r.id)

    return[r.toJSON() for r in rents]

    