package all

import (
	_ "github.com/aquasecurity/trivy/pkg/fanal/analyzer/buildinfo"
	_ "github.com/aquasecurity/trivy/pkg/fanal/analyzer/config/all"
	_ "github.com/aquasecurity/trivy/pkg/fanal/analyzer/executable"
	_ "github.com/aquasecurity/trivy/pkg/fanal/analyzer/imgconf/apk"
	_ "github.com/aquasecurity/trivy/pkg/fanal/analyzer/imgconf/dockerfile"
	_ "github.com/aquasecurity/trivy/pkg/fanal/analyzer/imgconf/secret"
	_ "github.com/aquasecurity/trivy/pkg/fanal/analyzer/language/c/conan"
	_ "github.com/aquasecurity/trivy/pkg/fanal/analyzer/language/conda/environment"
	_ "github.com/aquasecurity/trivy/pkg/fanal/analyzer/language/conda/meta"
	_ "github.com/aquasecurity/trivy/pkg/fanal/analyzer/language/dart/pub"
	_ "github.com/aquasecurity/trivy/pkg/fanal/analyzer/language/dotnet/deps"
	_ "github.com/aquasecurity/trivy/pkg/fanal/analyzer/language/dotnet/nuget"
	_ "github.com/aquasecurity/trivy/pkg/fanal/analyzer/language/dotnet/packagesprops"
	_ "github.com/aquasecurity/trivy/pkg/fanal/analyzer/language/elixir/mix"
	_ "github.com/aquasecurity/trivy/pkg/fanal/analyzer/language/golang/binary"
	_ "github.com/aquasecurity/trivy/pkg/fanal/analyzer/language/golang/mod"
	_ "github.com/aquasecurity/trivy/pkg/fanal/analyzer/language/java/gradle"
	_ "github.com/aquasecurity/trivy/pkg/fanal/analyzer/language/java/jar"
	_ "github.com/aquasecurity/trivy/pkg/fanal/analyzer/language/java/pom"
	_ "github.com/aquasecurity/trivy/pkg/fanal/analyzer/language/java/sbt"
	_ "github.com/aquasecurity/trivy/pkg/fanal/analyzer/language/julia/pkg"
	_ "github.com/aquasecurity/trivy/pkg/fanal/analyzer/language/nodejs/bun"
	_ "github.com/aquasecurity/trivy/pkg/fanal/analyzer/language/nodejs/npm"
	_ "github.com/aquasecurity/trivy/pkg/fanal/analyzer/language/nodejs/pkg"
	_ "github.com/aquasecurity/trivy/pkg/fanal/analyzer/language/nodejs/pnpm"
	_ "github.com/aquasecurity/trivy/pkg/fanal/analyzer/language/nodejs/yarn"
	_ "github.com/aquasecurity/trivy/pkg/fanal/analyzer/language/php/composer"
	_ "github.com/aquasecurity/trivy/pkg/fanal/analyzer/language/python/packaging"
	_ "github.com/aquasecurity/trivy/pkg/fanal/analyzer/language/python/pip"
	_ "github.com/aquasecurity/trivy/pkg/fanal/analyzer/language/python/pipenv"
	_ "github.com/aquasecurity/trivy/pkg/fanal/analyzer/language/python/poetry"
	_ "github.com/aquasecurity/trivy/pkg/fanal/analyzer/language/python/uv"
	_ "github.com/aquasecurity/trivy/pkg/fanal/analyzer/language/ruby/bundler"
	_ "github.com/aquasecurity/trivy/pkg/fanal/analyzer/language/ruby/gemspec"
	_ "github.com/aquasecurity/trivy/pkg/fanal/analyzer/language/rust/binary"
	_ "github.com/aquasecurity/trivy/pkg/fanal/analyzer/language/rust/cargo"
	_ "github.com/aquasecurity/trivy/pkg/fanal/analyzer/language/swift/cocoapods"
	_ "github.com/aquasecurity/trivy/pkg/fanal/analyzer/language/swift/swift"
	_ "github.com/aquasecurity/trivy/pkg/fanal/analyzer/licensing"
	_ "github.com/aquasecurity/trivy/pkg/fanal/analyzer/os/alpine"
	_ "github.com/aquasecurity/trivy/pkg/fanal/analyzer/os/amazonlinux"
	_ "github.com/aquasecurity/trivy/pkg/fanal/analyzer/os/debian"
	_ "github.com/aquasecurity/trivy/pkg/fanal/analyzer/os/redhatbase"
	_ "github.com/aquasecurity/trivy/pkg/fanal/analyzer/os/release"
	_ "github.com/aquasecurity/trivy/pkg/fanal/analyzer/os/ubuntu"
	_ "github.com/aquasecurity/trivy/pkg/fanal/analyzer/pkg/apk"
	_ "github.com/aquasecurity/trivy/pkg/fanal/analyzer/pkg/bottlerocket_inventory"
	_ "github.com/aquasecurity/trivy/pkg/fanal/analyzer/pkg/dpkg"
	_ "github.com/aquasecurity/trivy/pkg/fanal/analyzer/pkg/rpm"
	_ "github.com/aquasecurity/trivy/pkg/fanal/analyzer/repo/apk"
	_ "github.com/aquasecurity/trivy/pkg/fanal/analyzer/sbom"
	_ "github.com/aquasecurity/trivy/pkg/fanal/analyzer/secret"
)
