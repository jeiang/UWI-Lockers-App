{
  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixpkgs-unstable";

    devenv.url = "github:cachix/devenv";
    flake-parts = {
      url = "github:hercules-ci/flake-parts";
      inputs.nixpkgs-lib.follows = "nixpkgs";
    };
    treefmt-nix = {
      url = "github:numtide/treefmt-nix";
      inputs.nixpkgs.follows = "nixpkgs";
    };

    flake-compat = {
      url = "github:edolstra/flake-compat";
      flake = false;
    };
  };

  nixConfig = {
    extra-trusted-public-keys = "devenv.cachix.org-1:w1cLUi8dv3hnoSPGAuibQv+f9TZLr6cv/Hm9XgU50cw=";
    extra-substituters = "https://devenv.cachix.org";
  };

  outputs = {
    nixpkgs,
    devenv,
    flake-parts,
    treefmt-nix,
    ...
  } @ inputs:
    flake-parts.lib.mkFlake {inherit inputs;} {
      imports = [
        devenv.flakeModule
        treefmt-nix.flakeModule
      ];

      systems = nixpkgs.lib.systems.flakeExposed;

      perSystem = {
        config,
        pkgs,
        ...
      }: {
        # TODO: containers
        devenv.shells.default = {
          enterShell = ''
            set -euo pipefail

            # find the root of the project
            DIR=$PWD
            while
              RESULT=$(find "$DIR"/ -maxdepth 1 -name flake.nix)
              [[ -z $RESULT ]] && [[ "$DIR" != "/" ]]
            do DIR=$(dirname "$DIR"); done

            # Subshell so that original $PWD is not affected
            (
              cd "$DIR/frontend"
              ${pkgs.bun}/bin/bun install
            )
          '';
          packages = with pkgs; [
            bun
            commitizen
            config.treefmt.build.wrapper
            nodePackages.typescript-language-server
          ];

          languages.nix.enable = true;
          languages.javascript.enable = true;
          languages.go.enable = true;

          pre-commit.hooks.alejandra.enable = true;
          pre-commit.hooks.commitizen.enable = true;
          pre-commit.hooks.convco.enable = true;
          pre-commit.hooks.eslint.enable = true;
          pre-commit.hooks.treefmt.enable = true;
          pre-commit.hooks.gotest.enable = true;
          pre-commit.hooks.govet.enable = true;

          pre-commit.settings.treefmt.package = config.treefmt.build.wrapper;
          pre-commit.settings.eslint.extensions = "\.[jt]sx?$$";

          difftastic.enable = true;
        };

        treefmt = {
          projectRootFile = "flake.nix";
          programs = {
            alejandra.enable = true;
            deadnix.enable = true;
            prettier.enable = true;
            gofumpt.enable = true;
          };
        };
      };
    };
}
