{
  description = "GoUAE/golang.ae: The official site of the Go community in the UAE";

  outputs = inputs @ {flake-parts, ...}:
    flake-parts.lib.mkFlake {inherit inputs;} {
      imports = [{perSystem = {lib, ...}: {_module.args.l = lib // builtins;};}];
      systems = ["x86_64-linux" "aarch64-linux" "aarch64-darwin" "x86_64-darwin"];
      perSystem = {
        l,
        pkgs,
        config,
        inputs',
        ...
      }: let
        gomod2nix = inputs'.gomod2nix.legacyPackages;
      in {
        devShells.default = pkgs.mkShell {
          inputsFrom = [config.packages.default];

          packages = l.attrValues {
            inherit (pkgs) gopls;
            inherit (gomod2nix) gomod2nix;
          };
        };

        packages.default = gomod2nix.buildGoApplication {
          pname = "golang-ae";
          version = "0.1.0";

          pwd = ./.;
          src = ./.;

          modules = ./gomod2nix.toml;

          nativeBuildInputs = l.attrValues {
            inherit (pkgs) just;
            inherit (inputs'.templ.packages) templ;
          };

          preBuild = ''
            just gen
          '';

          meta.mainProgram = "golang.ae";
        };

        packages.image = pkgs.dockerTools.buildImage {
          name = "golang.ae";
          tag = config.packages.default.version;
          config.Cmd = ["${l.getExe config.packages.default}"];
        };
      };
    };

  inputs = {
    flake-parts.url = "github:hercules-ci/flake-parts";
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";

    templ = {
      url = "github:a-h/templ";
      inputs.nixpkgs.follows = "nixpkgs";
    };

    gomod2nix = {
      url = "github:nix-community/gomod2nix";
      inputs.nixpkgs.follows = "nixpkgs";
    };
  };
}
