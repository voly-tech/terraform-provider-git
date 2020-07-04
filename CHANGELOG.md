## 0.1.3 (July 04, 2020)

BUG FIXES:

* **Data Source** `git_repository`: get lightweight and annotated tags ([#39](https://github.com/innovationnorway/terraform-provider-git/issues/39))

## 0.1.2 (July 03, 2020)

BUG FIXES:

* **Data Source** `git_repository`: Fix nil pointer reference for remote `url` ([#36](https://github.com/innovationnorway/terraform-provider-git/issues/36))

## 0.1.1 (July 01, 2020)

ENHANCEMENTS:

* **Provider**: support for setting the `username`, `password`, `private_key`, `private_key_file`, `ignore_host_key` and `skip_tls_verify` properties ([#8](https://github.com/volcano-coffee-company/terraform-provider-git/issues/8), [#30](https://github.com/volcano-coffee-company/terraform-provider-git/issues/30))
* **Data Source** `git_repository`: support for setting the `branch`, `tag` and `url` properties ([#8](https://github.com/volcano-coffee-company/terraform-provider-git/issues/8))
* **Data Source** `git_repository`: support for setting default repository location using `GIT_DIR` environment variable ([#4](https://github.com/volcano-coffee-company/terraform-provider-git/issues/4))

## 0.1.0 (February 26, 2020)

- Initial release
