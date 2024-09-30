{
  description = "GoUAE/golang.ae: The official site of the GoUAE community.";

  outputs = inputs @ {
    flake-parts,
    treefmt,
    ...
  }:
    flake-parts.lib.mkFlake {inherit inputs;} {
      perSystem = {
        l,
        pkgs,
        config,
        inputs',
        ...
      }: {
        treefmt.config = {
          projectRootFile = "flake.nix";

          programs.gofmt.enable = true;
          programs.prettier.enable = true;
          programs.alejandra.enable = true;
        };

        devShells.default = pkgs.mkShell {
          packages = l.attrValues {
            inherit (pkgs) go just gopls;
            inherit (inputs'.templ.packages) templ;
            inherit (config.treefmt.build.programs) alejandra prettier;
          };
        };
      };

      imports = [treefmt.flakeModule {perSystem = {lib, ...}: {_module.args.l = lib // builtins;};}];
      systems = ["x86_64-linux" "aarch64-linux" "aarch64-darwin" "x86_64-darwin"];
    };

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    treefmt.url = "github:numtide/treefmt-nix";

    flake-parts.url = "github:hercules-ci/flake-parts";

    templ = {
      url = "github:a-h/templ";
      inputs.nixpkgs.follows = "nixpkgs";
    };
  };
}
