{
  description = "Go app with CGO and sqlite3 support";

  inputs = {
    nixpkgs.url = "nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = nixpkgs.legacyPackages.${system};
      in {
        packages.default = pkgs.buildGoModule {
          pname = "goNavigate";
          version = "0.1.0";

          src = ./.;

          vendorHash = "sha256-UXMr02pTdanneMe+X+KGPUYxPcHXLrAHuZ18l0Vdhic=";
          modRoot = ".";

          nativeBuildInputs = [ pkgs.gcc pkgs.sqlite ];
          env.CGO_ENABLED = "1";
          buildFlags = [ "-tags" "sqlite" ];
        };

        devShell = pkgs.mkShell {
          buildInputs = [ pkgs.go pkgs.gcc pkgs.sqlite ];

          CGO_ENABLED = "1";
          CC = "gcc";
        };
      }
    );
}

