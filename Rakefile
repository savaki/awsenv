require File.expand_path(File.dirname(__FILE__) + '/tasks/deb')

task :default => :package

task :clean do
  exec "rm -rf #{DIST}"
  exec "rm -f hosts/hosts"
end

namespace :go do 
  desc "go build"
  task :build do
    exec "(cd hosts ; go build)"
  end

  desc "go test"
  task :test do
    exec "go get -d -v ./..."
    exec "go test ./..."
  end
end

namespace :deb do 
  desc "package"
  task :package => :prepare do 
    create_package
  end

  desc "prepare"
  task :prepare => "go:build" do 
    prepare_package
  end
end
