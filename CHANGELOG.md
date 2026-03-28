# Changelog

## 0.4.0 (2026-03-28)

Full Changelog: [v0.3.1...v0.4.0](https://github.com/cjavdev/believe-cli/compare/v0.3.1...v0.4.0)

### Features

* add `--max-items` flag for paginated/streaming endpoints ([2bc0190](https://github.com/cjavdev/believe-cli/commit/2bc019097232419989b6cec6ea8239fd41e3ae80))
* add support for file downloads from binary response endpoints ([230380f](https://github.com/cjavdev/believe-cli/commit/230380fe1ce51ad0f0748888b2fac8e1564e3a20))
* Add ticket sales data model and more examples ([5c48fe6](https://github.com/cjavdev/believe-cli/commit/5c48fe6d6be471c674ce99ce77f522b229d80c54))
* **api:** manual updates ([8f9c1ce](https://github.com/cjavdev/believe-cli/commit/8f9c1cef193eed8211736c4a43810959e17f6fbd))
* **api:** manual updates ([a2ca735](https://github.com/cjavdev/believe-cli/commit/a2ca7350930c5a0548b8ec19c85734cb87af9f50))
* **api:** manual updates ([6a77c15](https://github.com/cjavdev/believe-cli/commit/6a77c15169463daf70bd85e42d3be0f9fcee02d8))
* **api:** manual updates ([916ad14](https://github.com/cjavdev/believe-cli/commit/916ad14d4efb677003664ea692f1d3461a43e1f6))
* **api:** manual updates ([f126e00](https://github.com/cjavdev/believe-cli/commit/f126e002c1aa7be74dc4ae27908d0759aee81580))
* **api:** manual updates ([c9d0336](https://github.com/cjavdev/believe-cli/commit/c9d0336d589b0342bf1b75ada04f5b606cc8a2b6))
* **api:** manual updates ([862ed5a](https://github.com/cjavdev/believe-cli/commit/862ed5af22baa149981e0d9ae8442aad56c68fd1))
* **api:** manual updates ([a76e244](https://github.com/cjavdev/believe-cli/commit/a76e2443035e5a676915b2e906439222771ef812))
* **api:** manual updates ([d0d6e47](https://github.com/cjavdev/believe-cli/commit/d0d6e47385e861ad05e6119ac1d841fa15a873f9))
* improved documentation and flags for client options ([affe157](https://github.com/cjavdev/believe-cli/commit/affe157582c273edd6755e68bbe3763f657fbda2))
* set CLI flag constant values automatically where `x-stainless-const` is set ([7463897](https://github.com/cjavdev/believe-cli/commit/7463897e8dd9892db4ad479a30ad5246157f3d49))
* support passing required body params through pipes ([d9de4bd](https://github.com/cjavdev/believe-cli/commit/d9de4bd19c0b93622edb8d463ea582041e896ff0))


### Bug Fixes

* add missing example parameters for test cases ([4e118c9](https://github.com/cjavdev/believe-cli/commit/4e118c9fa3fec9c25b611a51df4f951719446196))
* avoid printing usage errors twice ([4baad8f](https://github.com/cjavdev/believe-cli/commit/4baad8f8538a679f05d85eae65b9a4c28fee7176))
* avoid reading from stdin unless request body is form encoded or json ([83f69bb](https://github.com/cjavdev/believe-cli/commit/83f69bb74d599522e7e0a27184a368769b135f36))
* cli no longer hangs when stdin is attached to a pipe with empty input ([b14ce6d](https://github.com/cjavdev/believe-cli/commit/b14ce6dab63f356e010fd1d3436614b003752903))
* fix for encoding arrays with `any` type items ([cbf1fb1](https://github.com/cjavdev/believe-cli/commit/cbf1fb1002916a4b48eee9f4db6c1ff1e8c220ec))
* fix for off-by-one error in pagination logic ([7c75729](https://github.com/cjavdev/believe-cli/commit/7c75729b797d8285b6b6f506992ec6f34e2a5e2f))
* fix for test cases with newlines in YAML and better error reporting ([09a5761](https://github.com/cjavdev/believe-cli/commit/09a57618e9c8e21862a8b90102e8966144950b6f))
* improve linking behavior when developing on a branch not in the Go SDK ([bdee8ad](https://github.com/cjavdev/believe-cli/commit/bdee8adbb3d1b07214043d5d8b43c2e88428c414))
* more gracefully handle empty stdin input ([d8287c4](https://github.com/cjavdev/believe-cli/commit/d8287c42a804ee36b51d5d1437c250b2d8c956e1))
* no longer require an API key when building on production repos ([7524e67](https://github.com/cjavdev/believe-cli/commit/7524e6781d90b8629e42f6ae0d163630b896436f))
* pin formatting for headers to always use repeat/dot formats ([2ca74f3](https://github.com/cjavdev/believe-cli/commit/2ca74f3a0f4a3e7e4ef67cd80e3c9d8ce841fb5a))


### Chores

* **ci:** skip lint on metadata-only changes ([e4f2a77](https://github.com/cjavdev/believe-cli/commit/e4f2a7710796d3d8c8514ff95c4cf383296f08f9))
* **ci:** skip uploading artifacts on stainless-internal branches ([80d4472](https://github.com/cjavdev/believe-cli/commit/80d4472172268088e40152db4b1135bd38dafa5f))
* configure new SDK language ([50aedb3](https://github.com/cjavdev/believe-cli/commit/50aedb32aba44447789bb5c6ea7da383b8bd6b35))
* **internal:** codegen related update ([789c360](https://github.com/cjavdev/believe-cli/commit/789c3600402f2e32f5d249315b6e41e3775f1efb))
* **internal:** codegen related update ([e70f6e2](https://github.com/cjavdev/believe-cli/commit/e70f6e243befad5f542802fb6f96b924eb413e09))
* **internal:** codegen related update ([1d04885](https://github.com/cjavdev/believe-cli/commit/1d04885ce310f9443f6ebdaba6622bc681eaec18))
* **internal:** codegen related update ([3074c31](https://github.com/cjavdev/believe-cli/commit/3074c3116488dc1f628c158facccec7060f19270))
* **internal:** codegen related update ([f1595fb](https://github.com/cjavdev/believe-cli/commit/f1595fb9f4e1b313639cdf0417abffc7d6621033))
* **internal:** codegen related update ([b4721e3](https://github.com/cjavdev/believe-cli/commit/b4721e3faac1d589213d5ec1e1a6e49672e23d0f))
* **internal:** tweak CI branches ([0690096](https://github.com/cjavdev/believe-cli/commit/0690096bf3e0350f5a31f367f0cdd7ab1f3a5f96))
* **internal:** update gitignore ([f55da89](https://github.com/cjavdev/believe-cli/commit/f55da89fc6e5919cb7cb60ac31493fc09572332f))
* omit full usage information when missing required CLI parameters ([6c7160a](https://github.com/cjavdev/believe-cli/commit/6c7160a2e588a1b5ad030ac149d29f4432de3d16))
* update SDK settings ([a4218ea](https://github.com/cjavdev/believe-cli/commit/a4218eaebbee0a986477190f22445df01053dca1))
* zip READMEs as part of build artifact ([5c9d284](https://github.com/cjavdev/believe-cli/commit/5c9d28489deb0149d9b4a2e121ce5a26ea1f1e16))

## 0.3.1 (2026-03-12)

Full Changelog: [v0.3.0...v0.3.1](https://github.com/cjavdev/believe-cli/compare/v0.3.0...v0.3.1)

### Bug Fixes

* fix for test cases with newlines in YAML and better error reporting ([09a5761](https://github.com/cjavdev/believe-cli/commit/09a57618e9c8e21862a8b90102e8966144950b6f))

## 0.3.0 (2026-03-09)

Full Changelog: [v0.2.0...v0.3.0](https://github.com/cjavdev/believe-cli/compare/v0.2.0...v0.3.0)

### Features

* **api:** manual updates ([f126e00](https://github.com/cjavdev/believe-cli/commit/f126e002c1aa7be74dc4ae27908d0759aee81580))

## 0.2.0 (2026-03-09)

Full Changelog: [v0.1.0...v0.2.0](https://github.com/cjavdev/believe-cli/compare/v0.1.0...v0.2.0)

### Features

* **api:** manual updates ([c9d0336](https://github.com/cjavdev/believe-cli/commit/c9d0336d589b0342bf1b75ada04f5b606cc8a2b6))

## 0.1.0 (2026-03-09)

Full Changelog: [v0.0.1...v0.1.0](https://github.com/cjavdev/believe-cli/compare/v0.0.1...v0.1.0)

### Features

* add `--max-items` flag for paginated/streaming endpoints ([2bc0190](https://github.com/cjavdev/believe-cli/commit/2bc019097232419989b6cec6ea8239fd41e3ae80))
* add support for file downloads from binary response endpoints ([230380f](https://github.com/cjavdev/believe-cli/commit/230380fe1ce51ad0f0748888b2fac8e1564e3a20))
* Add ticket sales data model and more examples ([5c48fe6](https://github.com/cjavdev/believe-cli/commit/5c48fe6d6be471c674ce99ce77f522b229d80c54))
* **api:** manual updates ([862ed5a](https://github.com/cjavdev/believe-cli/commit/862ed5af22baa149981e0d9ae8442aad56c68fd1))
* **api:** manual updates ([a76e244](https://github.com/cjavdev/believe-cli/commit/a76e2443035e5a676915b2e906439222771ef812))
* **api:** manual updates ([d0d6e47](https://github.com/cjavdev/believe-cli/commit/d0d6e47385e861ad05e6119ac1d841fa15a873f9))
* improved documentation and flags for client options ([affe157](https://github.com/cjavdev/believe-cli/commit/affe157582c273edd6755e68bbe3763f657fbda2))
* support passing required body params through pipes ([d9de4bd](https://github.com/cjavdev/believe-cli/commit/d9de4bd19c0b93622edb8d463ea582041e896ff0))


### Bug Fixes

* add missing example parameters for test cases ([4e118c9](https://github.com/cjavdev/believe-cli/commit/4e118c9fa3fec9c25b611a51df4f951719446196))
* avoid printing usage errors twice ([4baad8f](https://github.com/cjavdev/believe-cli/commit/4baad8f8538a679f05d85eae65b9a4c28fee7176))
* fix for encoding arrays with `any` type items ([cbf1fb1](https://github.com/cjavdev/believe-cli/commit/cbf1fb1002916a4b48eee9f4db6c1ff1e8c220ec))
* more gracefully handle empty stdin input ([d8287c4](https://github.com/cjavdev/believe-cli/commit/d8287c42a804ee36b51d5d1437c250b2d8c956e1))
* pin formatting for headers to always use repeat/dot formats ([2ca74f3](https://github.com/cjavdev/believe-cli/commit/2ca74f3a0f4a3e7e4ef67cd80e3c9d8ce841fb5a))


### Chores

* **ci:** skip uploading artifacts on stainless-internal branches ([80d4472](https://github.com/cjavdev/believe-cli/commit/80d4472172268088e40152db4b1135bd38dafa5f))
* configure new SDK language ([50aedb3](https://github.com/cjavdev/believe-cli/commit/50aedb32aba44447789bb5c6ea7da383b8bd6b35))
* **internal:** codegen related update ([1d04885](https://github.com/cjavdev/believe-cli/commit/1d04885ce310f9443f6ebdaba6622bc681eaec18))
* **internal:** codegen related update ([3074c31](https://github.com/cjavdev/believe-cli/commit/3074c3116488dc1f628c158facccec7060f19270))
* **internal:** codegen related update ([f1595fb](https://github.com/cjavdev/believe-cli/commit/f1595fb9f4e1b313639cdf0417abffc7d6621033))
* **internal:** codegen related update ([b4721e3](https://github.com/cjavdev/believe-cli/commit/b4721e3faac1d589213d5ec1e1a6e49672e23d0f))
* update SDK settings ([a4218ea](https://github.com/cjavdev/believe-cli/commit/a4218eaebbee0a986477190f22445df01053dca1))
* zip READMEs as part of build artifact ([5c9d284](https://github.com/cjavdev/believe-cli/commit/5c9d28489deb0149d9b4a2e121ce5a26ea1f1e16))
