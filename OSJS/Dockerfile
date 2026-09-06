# Basic OS.js desktop build for MCS.OSJS (donor-derived neutral shell; no SCADA applications).
# Node 16 LTS: OS.js v3 uses webpack 4 (MD4), which OpenSSL 3 (Node 17+) rejects
# with ERR_OSSL_EVP_UNSUPPORTED. Node 16 (OpenSSL 1.1) builds cleanly.
FROM node:16-bullseye

WORKDIR /usr/src/osjs

# Install dependencies first (cached layer).
COPY package.json ./
RUN npm install --no-audit --no-fund

# Copy the desktop source.
COPY . .

# Local packages (theme/icons) have their own webpack build (osjs is an
# external provided by the core bundle). Build all local theme/icon packages
# before discovery so their dist/ exists, then discover (symlinks
# dist/apps|themes|icons/<Name> -> src/packages/<Name>/dist), then build the
# core client bundle. scripts/build-local-packages.js discovers every local
# package (src/packages/<Name>/webpack.config.js) and fails the build on any
# broken package, so discovery never points at a missing dist bundle.
RUN npm run build:local-packages
RUN npm run package:discover
RUN npm run build

# OSJS-003: the shell's management/UI port (see docs/NETWORK_EXPOSURE.md).
EXPOSE 18209
CMD ["npm", "run", "serve"]
