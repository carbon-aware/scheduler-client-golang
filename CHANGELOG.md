# Changelog

## 0.1.0-alpha.3 (2025-07-11)

Full Changelog: [v0.1.0-alpha.2...v0.1.0-alpha.3](https://github.com/carbon-aware/scheduler-client-golang/compare/v0.1.0-alpha.2...v0.1.0-alpha.3)

### Features

* **api:** api update ([696f44f](https://github.com/carbon-aware/scheduler-client-golang/commit/696f44ff25de2b13e9edae473b351d936eb7f530))
* **client:** add debug log helper ([c8c9064](https://github.com/carbon-aware/scheduler-client-golang/commit/c8c9064b5c90fa9e68d15c50d45a1bf3b2592fd0))
* **client:** add escape hatch for null slice & maps ([8bd6c73](https://github.com/carbon-aware/scheduler-client-golang/commit/8bd6c737ce1a69b339cf45972e4194d54b6e03f2))


### Bug Fixes

* don't try to deserialize as json when ResponseBodyInto is []byte ([459ede2](https://github.com/carbon-aware/scheduler-client-golang/commit/459ede25be56cc7a8bef8c2032449446f9dee7dc))


### Chores

* **ci:** enable for pull requests ([10f0892](https://github.com/carbon-aware/scheduler-client-golang/commit/10f0892a9f3c826deee121a1fbb9733f7438dcac))
* **ci:** only run for pushes and fork pull requests ([aa10145](https://github.com/carbon-aware/scheduler-client-golang/commit/aa1014513d21315b7cd0a0ce701ce055675e3136))
* fix documentation of null map ([45d2876](https://github.com/carbon-aware/scheduler-client-golang/commit/45d287660d873007119a6ef34c76ef080c66d33c))
* **internal:** fix lint script for tests ([6f0e397](https://github.com/carbon-aware/scheduler-client-golang/commit/6f0e397907fe52409623ab6e24a5d5a1660108fb))
* lint tests ([3bf9aa7](https://github.com/carbon-aware/scheduler-client-golang/commit/3bf9aa7248c2611ad5efb253009db2862d7f8fb5))
* lint tests in subpackages ([d6cb41e](https://github.com/carbon-aware/scheduler-client-golang/commit/d6cb41e392a27f5d094e64df7b5441d5f053a7c3))

## 0.1.0-alpha.2 (2025-06-04)

Full Changelog: [v0.1.0-alpha.1...v0.1.0-alpha.2](https://github.com/carbon-aware/scheduler-client-golang/compare/v0.1.0-alpha.1...v0.1.0-alpha.2)

### Features

* **api:** api update ([a374ebf](https://github.com/carbon-aware/scheduler-client-golang/commit/a374ebff63f19180d362f9c50a7f8bf01423e592))
* **api:** manual updates ([dc9cd38](https://github.com/carbon-aware/scheduler-client-golang/commit/dc9cd38df2f59f28520d1d811b1aa82665e15e6f))
* **client:** allow overriding unions ([004e060](https://github.com/carbon-aware/scheduler-client-golang/commit/004e06081d2bc2197e3cfd0b9a9627b69c786bea))


### Bug Fixes

* **client:** cast to raw message when converting to params ([88b6d4a](https://github.com/carbon-aware/scheduler-client-golang/commit/88b6d4ae7776d0366fd79cc9943f6bb0ce0cfe8f))
* **client:** correctly set stream key for multipart ([b86d274](https://github.com/carbon-aware/scheduler-client-golang/commit/b86d274d927e1f3caa1b9d1253adee795a912d6f))
* **client:** don't panic on marshal with extra null field ([a13bd69](https://github.com/carbon-aware/scheduler-client-golang/commit/a13bd699d94335ec530eaf2580f111d01706410e))
* fix error ([8d227c9](https://github.com/carbon-aware/scheduler-client-golang/commit/8d227c9a0c0c676eff89588cdee4b42c6991dcce))


### Chores

* **docs:** grammar improvements ([66a9190](https://github.com/carbon-aware/scheduler-client-golang/commit/66a91906b3cec06242610a58c0b28787e28de06d))
* improve devcontainer setup ([71407ea](https://github.com/carbon-aware/scheduler-client-golang/commit/71407ea6f1437c33c2cf3cfd307739863cf05968))
* make go mod tidy continue on error ([5a6324d](https://github.com/carbon-aware/scheduler-client-golang/commit/5a6324d2c542f60e0165af891fc814cd14dc7bc2))

## 0.1.0-alpha.1 (2025-05-13)

Full Changelog: [v0.0.1-alpha.0...v0.1.0-alpha.1](https://github.com/carbon-aware/scheduler-client-golang/compare/v0.0.1-alpha.0...v0.1.0-alpha.1)

### Features

* **api:** manual updates ([94aa316](https://github.com/carbon-aware/scheduler-client-golang/commit/94aa316c4b7a7750ce062bbc5d6080e7dd1e6a06))


### Chores

* configure new SDK language ([149e4c0](https://github.com/carbon-aware/scheduler-client-golang/commit/149e4c0f9518de2bb0d286203a6882696156e2d5))
* update SDK settings ([1218c86](https://github.com/carbon-aware/scheduler-client-golang/commit/1218c86a0032a8d29be5f93bbbaf7ac63c491687))
