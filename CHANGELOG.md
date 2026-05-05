# Changelog

<<<<<<< HEAD
=======
## 0.10.0 (2026-05-04)

Full Changelog: [v0.9.0...v0.10.0](https://github.com/cjavdev/believe-cli/compare/v0.9.0...v0.10.0)

### Features

* support passing path and query params over stdin ([eebf3b9](https://github.com/cjavdev/believe-cli/commit/eebf3b9ead1e0adaaa66ae0e2017b91cdc6a7a6d))


### Bug Fixes

* **cli:** correctly load zsh autocompletion ([475c1e9](https://github.com/cjavdev/believe-cli/commit/475c1e98771e2f797c0c424277cce864f9a020e9))
* flags for nullable body scalar fields are strictly typed ([7d33a49](https://github.com/cjavdev/believe-cli/commit/7d33a4937f6ca49ce4d842bc77c617399391df43))


### Chores

* **internal:** codegen related update ([c2a9c8d](https://github.com/cjavdev/believe-cli/commit/c2a9c8d4eb2011047cc9497810bbb35394f5b168))
* **internal:** codegen related update ([c1c300d](https://github.com/cjavdev/believe-cli/commit/c1c300dd6d986c87cde11f44ae4f348ca649aaab))

## 0.9.0 (2026-04-27)

Full Changelog: [v0.8.0...v0.9.0](https://github.com/cjavdev/believe-cli/compare/v0.8.0...v0.9.0)

### Features

* **api:** manual updates ([d48ea9d](https://github.com/cjavdev/believe-cli/commit/d48ea9d6db93da28f0da58956106529cb09a130c))

## 0.8.0 (2026-04-23)

Full Changelog: [v0.7.1...v0.8.0](https://github.com/cjavdev/believe-cli/compare/v0.7.1...v0.8.0)

### Features

* **api:** manual updates ([28daa51](https://github.com/cjavdev/believe-cli/commit/28daa51b1be7d7af5767f9787d5a85db94e51fd6))

## 0.7.1 (2026-04-23)

Full Changelog: [v0.7.0...v0.7.1](https://github.com/cjavdev/believe-cli/compare/v0.7.0...v0.7.1)

### Chores

* **internal:** codegen related update ([e0bccb4](https://github.com/cjavdev/believe-cli/commit/e0bccb47a918dae0bd951dc13dc963acd555bfd2))
* **internal:** more robust bootstrap script ([28c5d5c](https://github.com/cjavdev/believe-cli/commit/28c5d5cef968490fe3a05b41e6caea2c8f892508))

>>>>>>> 28c46ce (Apply custom code)
## 0.7.0 (2026-04-22)

Full Changelog: [v0.6.1...v0.7.0](https://github.com/cjavdev/believe-cli/compare/v0.6.1...v0.7.0)

### Features

* **api:** manual updates ([2a101a5](https://github.com/cjavdev/believe-cli/commit/2a101a52a7eaca5bddece92c32a2737802aa3b12))

## 0.6.1 (2026-04-20)

Full Changelog: [v0.6.0...v0.6.1](https://github.com/cjavdev/believe-cli/compare/v0.6.0...v0.6.1)

### Chores

* **internal:** codegen related update ([9105e7a](https://github.com/cjavdev/believe-cli/commit/9105e7a861631e78f9050e00d4a848fac325351e))

## 0.6.0 (2026-04-18)

Full Changelog: [v0.5.0...v0.6.0](https://github.com/cjavdev/believe-cli/compare/v0.5.0...v0.6.0)

### Features

* **cli:** add `--raw-output`/`-r` option to print raw (non-JSON) strings ([60278a5](https://github.com/cjavdev/believe-cli/commit/60278a56a1fad531964f65c8cd84cd11271141fe))
* **cli:** alias parameters in data with `x-stainless-cli-data-alias` ([5ba7273](https://github.com/cjavdev/believe-cli/commit/5ba72737bec644bf34adaff5361e738c7f6a45b8))
* **cli:** send filename and content type when reading input from files ([aefad6b](https://github.com/cjavdev/believe-cli/commit/aefad6b8bd3a80018b6a81f496226744edae09b3))


### Chores

* add documentation for ./scripts/link ([aa92acc](https://github.com/cjavdev/believe-cli/commit/aa92accef5b780aae606ed6a8ba280ea7f7fcefc))
* **ci:** support manually triggering release workflow ([b1ac39e](https://github.com/cjavdev/believe-cli/commit/b1ac39ee7adcd83cdb257cec5a8bc309b232a91a))
* **cli:** fall back to JSON when using default "explore" with non-TTY ([31024cc](https://github.com/cjavdev/believe-cli/commit/31024ccf6d29c8c33761bdb049b91592cd620529))
* **cli:** switch long lists of positional args over to param structs ([e221d5a](https://github.com/cjavdev/believe-cli/commit/e221d5a129243c26e4206c8ce3a197cab49fb19b))
* **cli:** use `ShowJSONOpts` as argument to `formatJSON` instead of many positionals ([4355bad](https://github.com/cjavdev/believe-cli/commit/4355bad4b33a6eb8fd9728b3112dc988f3c85158))

## 0.5.0 (2026-04-13)

Full Changelog: [v0.4.1...v0.5.0](https://github.com/cjavdev/believe-cli/compare/v0.4.1...v0.5.0)

### Features

* **api:** release PRs ([1527e97](https://github.com/cjavdev/believe-cli/commit/1527e97a4cd48cdb533b4726473ee8bf145570d9))

## 0.4.1 (2026-04-10)

Full Changelog: [v0.4.0...v0.4.1](https://github.com/cjavdev/believe-cli/compare/v0.4.0...v0.4.1)

### Bug Fixes

* fall back to main branch if linking fails in CI ([8ac0e4d](https://github.com/cjavdev/believe-cli/commit/8ac0e4dd27b239c3170834a634230ea9647c38f2))
* fix for failing to drop invalid module replace in link script ([8170ff3](https://github.com/cjavdev/believe-cli/commit/8170ff3c49b6b84fbc0358aba125adbbe4b7dda3))
* fix quoting typo ([ba21723](https://github.com/cjavdev/believe-cli/commit/ba21723f12726d87b22817c08a200c0620f0c07c))


### Chores

* **cli:** additional test cases for `ShowJSONIterator` ([3388c47](https://github.com/cjavdev/believe-cli/commit/3388c476117ff686398c936f805d049aa3c0dbd4))
* **cli:** let `--format raw` be used in conjunction with `--transform` ([c9e47fd](https://github.com/cjavdev/believe-cli/commit/c9e47fda8cc5a9bf39a4fa1b697f743c996e80c8))
* **internal:** codegen related update ([6e4d314](https://github.com/cjavdev/believe-cli/commit/6e4d3140a0a6085ccea2ad9be12e0740cc3bbbd3))
* modify CLI tests to inject stdout so mutating `os.Stdout` isn't necessary ([8b0d57d](https://github.com/cjavdev/believe-cli/commit/8b0d57d69fde26668f886a337121d0c4c287e045))

## 0.4.0 (2026-04-03)

Full Changelog: [v0.3.1...v0.4.0](https://github.com/cjavdev/believe-cli/compare/v0.3.1...v0.4.0)

### Features

* add `--max-items` flag for paginated/streaming endpoints ([2bc0190](https://github.com/cjavdev/believe-cli/commit/2bc019097232419989b6cec6ea8239fd41e3ae80))
* add support for file downloads from binary response endpoints ([230380f](https://github.com/cjavdev/believe-cli/commit/230380fe1ce51ad0f0748888b2fac8e1564e3a20))
* Add ticket sales data model and more examples ([5c48fe6](https://github.com/cjavdev/believe-cli/commit/5c48fe6d6be471c674ce99ce77f522b229d80c54))
* allow `-` as value representing stdin to binary-only file parameters in CLIs ([4a4a641](https://github.com/cjavdev/believe-cli/commit/4a4a64108a212e4955414e3764f7427ee3cda382))
* **api:** manual updates ([8f9c1ce](https://github.com/cjavdev/believe-cli/commit/8f9c1cef193eed8211736c4a43810959e17f6fbd))
* **api:** manual updates ([a2ca735](https://github.com/cjavdev/believe-cli/commit/a2ca7350930c5a0548b8ec19c85734cb87af9f50))
* **api:** manual updates ([6a77c15](https://github.com/cjavdev/believe-cli/commit/6a77c15169463daf70bd85e42d3be0f9fcee02d8))
* **api:** manual updates ([916ad14](https://github.com/cjavdev/believe-cli/commit/916ad14d4efb677003664ea692f1d3461a43e1f6))
* **api:** manual updates ([f126e00](https://github.com/cjavdev/believe-cli/commit/f126e002c1aa7be74dc4ae27908d0759aee81580))
* **api:** manual updates ([c9d0336](https://github.com/cjavdev/believe-cli/commit/c9d0336d589b0342bf1b75ada04f5b606cc8a2b6))
* **api:** manual updates ([862ed5a](https://github.com/cjavdev/believe-cli/commit/862ed5af22baa149981e0d9ae8442aad56c68fd1))
* **api:** manual updates ([a76e244](https://github.com/cjavdev/believe-cli/commit/a76e2443035e5a676915b2e906439222771ef812))
* **api:** manual updates ([d0d6e47](https://github.com/cjavdev/believe-cli/commit/d0d6e47385e861ad05e6119ac1d841fa15a873f9))
* better error message if scheme forgotten in CLI `*_BASE_URL`/`--base-url` ([f3e3707](https://github.com/cjavdev/believe-cli/commit/f3e3707f684537c08924ac05a6aad5d48df3750c))
* binary-only parameters become CLI flags that take filenames only ([c80b0a6](https://github.com/cjavdev/believe-cli/commit/c80b0a659d309a7579fa49ac091f12fb6483e0cc))
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
* handle empty data set using `--format explore` ([2723724](https://github.com/cjavdev/believe-cli/commit/2723724d9e7b9ee87252d7c2cf899acd6cbd2fff))
* improve linking behavior when developing on a branch not in the Go SDK ([bdee8ad](https://github.com/cjavdev/believe-cli/commit/bdee8adbb3d1b07214043d5d8b43c2e88428c414))
* more gracefully handle empty stdin input ([d8287c4](https://github.com/cjavdev/believe-cli/commit/d8287c42a804ee36b51d5d1437c250b2d8c956e1))
* no longer require an API key when building on production repos ([7524e67](https://github.com/cjavdev/believe-cli/commit/7524e6781d90b8629e42f6ae0d163630b896436f))
* pin formatting for headers to always use repeat/dot formats ([2ca74f3](https://github.com/cjavdev/believe-cli/commit/2ca74f3a0f4a3e7e4ef67cd80e3c9d8ce841fb5a))
* use `RawJSON` when iterating items with `--format explore` in the CLI ([f2d73a8](https://github.com/cjavdev/believe-cli/commit/f2d73a83f3487cd6a0370d54a047c6f7a8bfaa73))


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
* mark all CLI-related tests in Go with `t.Parallel()` ([6884a13](https://github.com/cjavdev/believe-cli/commit/6884a13b31911c120de1c01028ed290762754bbd))
* omit full usage information when missing required CLI parameters ([6c7160a](https://github.com/cjavdev/believe-cli/commit/6c7160a2e588a1b5ad030ac149d29f4432de3d16))
* switch some CLI Go tests from `os.Chdir` to `t.Chdir` ([3f5682f](https://github.com/cjavdev/believe-cli/commit/3f5682f33bff5e3d26b767c355b3c313f1f19e77))
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
