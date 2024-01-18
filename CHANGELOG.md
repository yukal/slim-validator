## [v0.21.0](https://github.com/yukal/slim-validator/compare/6e59d0cff354da1dff447271fd1eb6b4db93d115...e1fcbe96de9320c56419775f15d749c8bab1b6b2) – 2024-01-18

### New Features

- validation type improvement
  - **check-others:**  ([e36919cb](https://github.com/yukal/slim-validator/commit/e36919cb8edab37a4f12af1385d021147c430d60))
  - **check-field:**  ([3dbf0460](https://github.com/yukal/slim-validator/commit/3dbf04602bcdb3a47db6e9e64efa3cd1906e6c7d))
- return MsgInvalidValue
  - **filter-range:**  ([236da085](https://github.com/yukal/slim-validator/commit/236da085799cb3d11a2ffb2f79d32dcb18c39da9))
  - **filter-eq:**  ([c83bae1e](https://github.com/yukal/slim-validator/commit/c83bae1e89c57d36bcae4446e54bc6cc48bd569e))
  - **filter-max:**  ([b53b7e42](https://github.com/yukal/slim-validator/commit/b53b7e42124ec403ab5dd0b162ff96ed8adfc119))
  - **filter-min:**  ([2e7c907f](https://github.com/yukal/slim-validator/commit/2e7c907fb6ef7b16e7983c555354523a65656a20))
- return MsgInvalidRule
  - **modifier-each:**  ([9a60791a](https://github.com/yukal/slim-validator/commit/9a60791a4535421d63a7e0170f8439ff00accd16))

### Refactors

- **core:** remove unnecessary check
  - **filter:**  ([47523c60](https://github.com/yukal/slim-validator/commit/47523c608fee2cba630f28c178cfdea1b61dba1c))
  - **compare:**  ([4fab961d](https://github.com/yukal/slim-validator/commit/4fab961dc7f553b73033cf13b6e1e3c890699025))
- **core:** func IsYearEqual became a private filterYearEqual ([ee9f1d0a](https://github.com/yukal/slim-validator/commit/ee9f1d0a53930dbe3fde3912363ea01d55894469))
- **test:**  remove returns that have the same default value ([3573d9c0](https://github.com/yukal/slim-validator/commit/3573d9c06e431002e71b69e097e3b2a6a2a94e54))
- **test:**  combine logical blocks ([8f061876](https://github.com/yukal/slim-validator/commit/8f061876f46a320fa5e3d7f3ff46660293978f77))

### Tests

- **core:**  add benchmark tests ([6d739899](https://github.com/yukal/slim-validator/commit/6d739899b632098294287e6c99cfa34f80d31377))
- **core:**  check invalidity within missed tests ([a83bed3a](https://github.com/yukal/slim-validator/commit/a83bed3a1973ee1610a6013abd7bc8bae57c7c28))
- **core:**  validate optional fields ([cf7011af](https://github.com/yukal/slim-validator/commit/cf7011af12027a80fbb57043d7413b71a06cb629))
- **coverage:**  improve coverage indicators ([b861f45c](https://github.com/yukal/slim-validator/commit/b861f45caa88b7b9e0952e4d6d4a143f1e2bc64c))
-  skip test files of private methods ([5ac2547c](https://github.com/yukal/slim-validator/commit/5ac2547c7f003044646cb2c132e69c0b214878fe))
- **compare:**  remove duplicate tests ([8b681185](https://github.com/yukal/slim-validator/commit/8b681185074dfda1f2b141ca8f9109034e933c11))


### Build System

-  add Makefile ([7d0ba005](https://github.com/yukal/slim-validator/commit/7d0ba00583218ff4c1452b6dfdc859668f0f09eb))

### Continuous Integration

- **job-action:**  turn on benchmark tests ([05f4dc94](https://github.com/yukal/slim-validator/commit/05f4dc944074b24bab3e7c3db9914a5e7a064dc0))

### Docs

-  update readme ([e1fcbe96](https://github.com/yukal/slim-validator/commit/e1fcbe96de9320c56419775f15d749c8bab1b6b2))


## [v0.14.0](https://github.com/yukal/slim-validator/compare/dfdb22c433879ea3b12f5727f3c349148465d61c...656a76596b01860c61dc8213f35ff62a949117dc) – 2024-01-15

### New Features

- **core:**  validate time ([7c398586](https://github.com/yukal/slim-validator/commit/7c3985863f7a6d1a36926e755eb70d98572c2e53))
- **core:**  validate date ([90c31cc7](https://github.com/yukal/slim-validator/commit/90c31cc7d942f3162de72885068fc9eba665d50e))

### Refactors

- **core:**  update constants ([f025490b](https://github.com/yukal/slim-validator/commit/f025490b5662601de5c0de7f7b354ceb0b5d6e79))

### Tests

- **core:**  modifier time ([656a7659](https://github.com/yukal/slim-validator/commit/656a76596b01860c61dc8213f35ff62a949117dc))
- **core:**  modifier date ([6ad7a758](https://github.com/yukal/slim-validator/commit/6ad7a7584630f9ef5d4356dc391ade90d67ffe41))

### Docs

- **readme:**  the time modifier description ([032021ea](https://github.com/yukal/slim-validator/commit/032021ea62225bb1930e09866303a48bc66f7009))
- **readme:**  the date modifier description ([320858e8](https://github.com/yukal/slim-validator/commit/320858e8db61b84f955d1ab731d87526c6d25465))


## [v0.12.0](https://github.com/yukal/slim-validator/compare/87e27661ecc321bc6b45ec477d1cb5c80f63fbfd...719673093ca141fadcd9ec18115b0673fba00f46) – 2024-01-13

### New Features

- **core:**  add the modifier "each" ([c62d0b59](https://github.com/yukal/slim-validator/commit/c62d0b5970314081344028c1b076185d7c37ce57))
- **core:**  add the modifier "fields" ([54903f64](https://github.com/yukal/slim-validator/commit/54903f647c5ee41ba8d767fa01c957e6bcc62c9a))
- **core:**  implement isValid method ([638e303a](https://github.com/yukal/slim-validator/commit/638e303a58cbd5f845c85ee3ac10936518a62043))
- **core:**  validate min ([6d5a0a0e](https://github.com/yukal/slim-validator/commit/6d5a0a0efa0e2fdb2800da6e7225f8a1c74e9b80))
- **core:**  validate max ([0d0004ed](https://github.com/yukal/slim-validator/commit/0d0004ed0c54e6aca89586c7d6a3a744408991fc))
- **core:**  validator.Group ([1f960087](https://github.com/yukal/slim-validator/commit/1f960087c88659dfd5f51e080e50cf3e4ef83aa7))
- **core:**  validate eq ([06a38315](https://github.com/yukal/slim-validator/commit/06a38315da14aa1e770ad9f0c58dba007f9b0ecc))
- **core:**  validate range ([0cbd4411](https://github.com/yukal/slim-validator/commit/0cbd441189fc23be12a0a751d1c030519c977054))
- **core:**  validate year ([e5592037](https://github.com/yukal/slim-validator/commit/e559203799c0e04f44e7863b45d70a4266c3bc05))
- **core:**  validate non-zero ([3645545e](https://github.com/yukal/slim-validator/commit/3645545e09f34e7ca9915ea35556a3cd45393751))
- **core:**  validate match ([e88c83d7](https://github.com/yukal/slim-validator/commit/e88c83d784c4f6cb11a58ce7bd78aad6bb52ab03))
- **core:**  validate each-match ([86a3531b](https://github.com/yukal/slim-validator/commit/86a3531bbdfc299dfa7ca2b7f5877f9f29576ff1))

### Refactors

- **filter-match:**  update logic and tests ([3220500e](https://github.com/yukal/slim-validator/commit/3220500e53e12141fab4bc8202efd828f7cea8e5))
- **filter-match:**  rename IsMatch method ([bc4392c0](https://github.com/yukal/slim-validator/commit/bc4392c0f7931fc750f31eedb1a964f164edb79c))
- **each-match:**  remove deprecated eachMatch ([d32baff5](https://github.com/yukal/slim-validator/commit/d32baff5fa80a3dc3e2dd399fccb01f43d096bba))
- **compare:**  invalid rule ([3684f9e8](https://github.com/yukal/slim-validator/commit/3684f9e83867b5304ebb6b9b070cca4d0d9b3777))
- **range:**  support numeric types ([f01560b8](https://github.com/yukal/slim-validator/commit/f01560b8ffbd780dc7de4ae862acae3bf754f293))
- **check-field:**  recursively unpack interface value ([65e7c931](https://github.com/yukal/slim-validator/commit/65e7c931530fc5e9103fe00f7d9431ed7bdfeff6))

### Tests

- **core:**  modifier each-match ([dda43810](https://github.com/yukal/slim-validator/commit/dda4381050c49a85c68fc49b2e96c7ca66f9ba59))
- **core:**  modifier each-min ([8c9b4507](https://github.com/yukal/slim-validator/commit/8c9b450760aa0cb41b9b4f2f9e8951c7f0c26c57))
- **core:**  modifier each-max ([5bfe9c84](https://github.com/yukal/slim-validator/commit/5bfe9c84566e7fb403677bfbc87558cff65c80ed))
- **core:**  modifier each-eq ([99255020](https://github.com/yukal/slim-validator/commit/992550207d7d17ea3f92189465fafadf72bfd26f))
- **core:**  modifier each-range ([833143f0](https://github.com/yukal/slim-validator/commit/833143f00e1cc9f19079cf9addcefab2641afaee))
- **core:**  modifier fields ([b194f206](https://github.com/yukal/slim-validator/commit/b194f206d8a9c5bc5758ea269576eb3360ee186a))
- **core:**  isValid ([3e0681f5](https://github.com/yukal/slim-validator/commit/3e0681f5dc107ac4cb6eee0b361b8889e649ed11))
- **core:**  validate min ([90fac12e](https://github.com/yukal/slim-validator/commit/90fac12eaa35206f60da6b7d4a2c4fc51519b50b))
- **core:**  validate max ([a1e5057c](https://github.com/yukal/slim-validator/commit/a1e5057c68adc09f0833a9b7a7ed4d69c84152be))
- **core:**  validator.Group ([6583a86f](https://github.com/yukal/slim-validator/commit/6583a86f77c02f8f9bdc4ccddc3d3ea55fc12fb0))
- **core:**  validate eq ([779b8b95](https://github.com/yukal/slim-validator/commit/779b8b951d2588fa1d6ebd2db5825840877ded11))
- **core:**  validate range ([5f2faa2b](https://github.com/yukal/slim-validator/commit/5f2faa2bcca2ea3822c27e76ee719c8653b73e87))
- **core:**  validate year ([ea0ed671](https://github.com/yukal/slim-validator/commit/ea0ed67148bbc430c99539f57c98836e461a511a))
- **core:**  validate non-zero ([affebf06](https://github.com/yukal/slim-validator/commit/affebf063a75a1446355d5107145cb017dd0fc23))
- **core:**  validate match ([c085f2f4](https://github.com/yukal/slim-validator/commit/c085f2f4a0450ce03c74503a29be3dc074d441cc))
- **core:**  validate each-match ([141dc870](https://github.com/yukal/slim-validator/commit/141dc870603bf6fd7c6a03e908ea3a94c1785ce8))

### Chores

- **core:**  improve descriptions ([b018153e](https://github.com/yukal/slim-validator/commit/b018153e50b065e92315e89734b2e3f69bc5b8d0))
- **changelog:**  improve changelog reading style ([50a7b554](https://github.com/yukal/slim-validator/commit/50a7b55415c633c8e4eed4638530a39e626cbe2d))
- **changelog:**  add workflow ([16258ef8](https://github.com/yukal/slim-validator/commit/16258ef8edfa7a93287dfa59037c71abe77e40a1))
- **changelog:**  add workflow ([4f5ab236](https://github.com/yukal/slim-validator/commit/4f5ab23697d5fb8c6c77ff0851c83b967f60a4d9)) ([a80a39e3](https://github.com/yukal/slim-validator/commit/a80a39e3fc68bdf0ab6cf7971838d6d567252c8b)) ([b84081d2](https://github.com/yukal/slim-validator/commit/b84081d24f31ff04170dab1295537cc2bf883653)) ([1c02384a](https://github.com/yukal/slim-validator/commit/1c02384a2c2e4fad46718e1769083a297a518e01))
- **init:**  add goblin ([508a21b3](https://github.com/yukal/slim-validator/commit/508a21b37269cf999d0527f03af20e0b19a34f0e))
- **init:**  add license   ([87e2766](https://github.com/yukal/slim-validator/commit/87e27661ecc321bc6b45ec477d1cb5c80f63fbfd))
- **init:**  add README.md ([c62f9bd](https://github.com/yukal/slim-validator/commit/c62f9bd7dfc3be790f3f7bc7315e40043c1a5513))
- **init:**  add go module ([4494ca4](https://github.com/yukal/slim-validator/commit/4494ca40ce14d41e7c4ca778e2979df86150e9ba))

### Docs

- **readme:**  improve readme ([71967309](https://github.com/yukal/slim-validator/commit/719673093ca141fadcd9ec18115b0673fba00f46))
- **readme:**  add links ([ef80a9ff](https://github.com/yukal/slim-validator/commit/ef80a9ff5e190ccb9db7f99317fd97a36788e05a))
- **readme:**  add description with examples
  - min ([8a9a6d34](https://github.com/yukal/slim-validator/commit/8a9a6d349113a624638675d572d3056c1b8340bf)) ([08826219](https://github.com/yukal/slim-validator/commit/08826219e08a5f2e1033d942635995a8ea8cec1f))
  - max ([c6bbf40a](https://github.com/yukal/slim-validator/commit/c6bbf40ac2e6e42a61f962f0339f2c47d5fdeddc))
  - eq ([628c9c26](https://github.com/yukal/slim-validator/commit/628c9c26a96b103459af615e989b7ccc6a35971e))
  - range  ([85d6e48a](https://github.com/yukal/slim-validator/commit/85d6e48a2c2a82f68961a118623833753bb41209))
  - year ([a0809815](https://github.com/yukal/slim-validator/commit/a08098150df4149189ad564d4740c844d971ed47))
- **readme:**  add description with examples ([167b0bc2](https://github.com/yukal/slim-validator/commit/167b0bc233ca865691df94e3921295697b39fe1b)) ([e9cfce77](https://github.com/yukal/slim-validator/commit/e9cfce77ac2b142526144be7112a62bdbff18483))

### Continuous Integration

- **action:**  back to non-nested single runners ([d45cb7e0](https://github.com/yukal/slim-validator/commit/d45cb7e0cec2bb1ffa3217bed9cde6aa8201369e)) ([161bce71](https://github.com/yukal/slim-validator/commit/161bce710ed2db15d0efe28dbcd5c3561c24dae4))
- **security:**  add dependabot ([42b358ed](https://github.com/yukal/slim-validator/commit/42b358ed725daf31c34916d22e1429ff3ba82291)) ([291e5bf3](https://github.com/yukal/slim-validator/commit/291e5bf31c97cac212d5f5ca47734afeecae55df))
- **action:**  add test job ([3f6e18b6](https://github.com/yukal/slim-validator/commit/3f6e18b6c36eb3256535168ece2c26a14a17d651))
- **action:**  add lint job ([47bc152e](https://github.com/yukal/slim-validator/commit/47bc152ec3216b04ac27983fc2a6fcb2d815e16a))
- **action:**  add release job ([456ddcb6](https://github.com/yukal/slim-validator/commit/456ddcb67c3271efe2e9c2e11feb5a8d3a744e63))
- **action:**  add develop workflow ([db391bb0](https://github.com/yukal/slim-validator/commit/db391bb0397d95705ef4b22bd24a826165d4aca1))
- **action:**  add release workflow ([9662b8ec](https://github.com/yukal/slim-validator/commit/9662b8ec1e25fdef38bfd10418d4a0b93c50c059))
- **release:**  add semantic release config ([b09be68b](https://github.com/yukal/slim-validator/commit/b09be68b53bd2888baa3ae1087f5e86f87de0574))
