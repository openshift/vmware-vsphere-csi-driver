# Release Process

<<<<<<< HEAD
||||||| parent of 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
## Semantic Convention Generation

New versions of the [OpenTelemetry specification] mean new versions of the `semconv` package need to be generated.
The `semconv-generate` make target is used for this.

1. Checkout a local copy of the [OpenTelemetry specification] to the desired release tag.
2. Run the `make semconv-generate ...` target from this repository.

For example,

```sh
export TAG="v1.7.0" # Change to the release version you are generating.
export OTEL_SPEC_REPO="/absolute/path/to/opentelemetry-specification"
git -C "$OTEL_SPEC_REPO" checkout "tags/$TAG"
make semconv-generate # Uses the exported TAG and OTEL_SPEC_REPO.
```

This should create a new sub-package of [`semconv`](./semconv).
Ensure things look correct before submitting a pull request to include the addition.

=======
## Semantic Convention Generation

New versions of the [OpenTelemetry Semantic Conventions] mean new versions of the `semconv` package need to be generated.
The `semconv-generate` make target is used for this.

1. Checkout a local copy of the [OpenTelemetry Semantic Conventions] to the desired release tag.
2. Pull the latest `otel/semconvgen` image: `docker pull otel/semconvgen:latest`
3. Run the `make semconv-generate ...` target from this repository.

For example,

```sh
export TAG="v1.21.0" # Change to the release version you are generating.
export OTEL_SEMCONV_REPO="/absolute/path/to/opentelemetry/semantic-conventions"
docker pull otel/semconvgen:latest
make semconv-generate # Uses the exported TAG and OTEL_SEMCONV_REPO.
```

This should create a new sub-package of [`semconv`](./semconv).
Ensure things look correct before submitting a pull request to include the addition.

## Breaking changes validation

You can run `make gorelease` that runs [gorelease](https://pkg.go.dev/golang.org/x/exp/cmd/gorelease) to ensure that there are no unwanted changes done in the public API.

You can check/report problems with `gorelease` [here](https://golang.org/issues/26420).

>>>>>>> 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
## Pre-Release

Update go.mod for submodules to depend on the new release which will happen in the next step.

1. Run the pre-release script. It creates a branch `pre_release_<new tag>` that will contain all release changes.

    ```
    ./pre_release.sh -t <new tag>
    ```

2. Verify the changes.

    ```
    git diff main
    ```

    This should have changed the version for all modules to be `<new tag>`.

3. Update the [Changelog](./CHANGELOG.md).
   - Make sure all relevant changes for this release are included and are in language that non-contributors to the project can understand.
       To verify this, you can look directly at the commits since the `<last tag>`.

       ```
       git --no-pager log --pretty=oneline "<last tag>..HEAD"
       ```

   - Move all the `Unreleased` changes into a new section following the title scheme (`[<new tag>] - <date of release>`).
   - Update all the appropriate links at the bottom.

4. Push the changes to upstream and create a Pull Request on GitHub.
    Be sure to include the curated changes from the [Changelog](./CHANGELOG.md) in the description.


## Tag

Once the Pull Request with all the version changes has been approved and merged it is time to tag the merged commit.

***IMPORTANT***: It is critical you use the same tag that you used in the Pre-Release step!
Failure to do so will leave things in a broken state.

***IMPORTANT***: [There is currently no way to remove an incorrectly tagged version of a Go module](https://github.com/golang/go/issues/34189).
It is critical you make sure the version you push upstream is correct.
[Failure to do so will lead to minor emergencies and tough to work around](https://github.com/open-telemetry/opentelemetry-go/issues/331).

1. Run the tag.sh script using the `<commit-hash>` of the commit on the main branch for the merged Pull Request.

    ```
    ./tag.sh <new tag> <commit-hash>
    ```

2. Push tags to the upstream remote (not your fork: `github.com/open-telemetry/opentelemetry-go.git`).
    Make sure you push all sub-modules as well.

    ```
    git push upstream <new tag>
    git push upstream <submodules-path/new tag>
    ...
    ```

## Release

Finally create a Release for the new `<new tag>` on GitHub.
The release body should include all the release notes from the Changelog for this release.
Additionally, the `tag.sh` script generates commit logs since last release which can be used to supplement the release notes.

## Verify Examples

After releasing verify that examples build outside of the repository.

```
./verify_examples.sh
```

The script copies examples into a different directory removes any `replace` declarations in `go.mod` and builds them.
This ensures they build with the published release, not the local copy.

## Contrib Repository

Once verified be sure to [make a release for the `contrib` repository](https://github.com/open-telemetry/opentelemetry-go-contrib/blob/main/RELEASING.md) that uses this release.
<<<<<<< HEAD
||||||| parent of 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))

### Website Documentation

Update [the documentation](./website_docs) for [the OpenTelemetry website](https://opentelemetry.io/docs/go/).
Importantly, bump any package versions referenced to be the latest one you just released and ensure all code examples still compile and are accurate.

[OpenTelemetry specification]: https://github.com/open-telemetry/opentelemetry-specification
=======

### Website Documentation

Update the [Go instrumentation documentation] in the OpenTelemetry website under [content/en/docs/instrumentation/go].
Importantly, bump any package versions referenced to be the latest one you just released and ensure all code examples still compile and are accurate.

[OpenTelemetry Semantic Conventions]: https://github.com/open-telemetry/semantic-conventions
[Go instrumentation documentation]: https://opentelemetry.io/docs/instrumentation/go/
[content/en/docs/instrumentation/go]: https://github.com/open-telemetry/opentelemetry.io/tree/main/content/en/docs/instrumentation/go

### Demo Repository

Bump the dependencies in the following Go services:

- [`accountingservice`](https://github.com/open-telemetry/opentelemetry-demo/tree/main/src/accountingservice)
- [`checkoutservice`](https://github.com/open-telemetry/opentelemetry-demo/tree/main/src/checkoutservice)
- [`productcatalogservice`](https://github.com/open-telemetry/opentelemetry-demo/tree/main/src/productcatalogservice)
>>>>>>> 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
