document.addEventListener('DOMContentLoaded', function () {

    function keep_alive_server() {
     fetch ("/flaskwebgui-dumb-request-for-middleware-keeping-the-server-online", {
        method: 'GET',
        cache: 'no-cache'
      })
        .then(res => { })
        .catch(err => { })
    }
    
    try {
      setInterval(keep_alive_server, 2 * 1000)()
    } catch (error) {
      // doesn't matter handled by middleware
    }
    
    })
    

