desc "Build chromedriver_helper to bin/"
task :build do
  sh "go build -o bin/chromedriver_helper"
end

desc "Run test"
task :test do
  src_path = File.join(ENV["GOPATH"], "src")
  path = File.expand_path("../...", __FILE__).sub("#{src_path}/", "")
  sh "go test #{path}"
end

task :default => :test
