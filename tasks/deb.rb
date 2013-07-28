DIST    = "dist"
VERSION = "1.0"

def create_package
  system <<EOF
  # use fpm to package the scripts
  fpm \
    --force \
    --deb-user 0 \
    --deb-group 0 \
    --url http://github.com/savaki/awsenv/hosts \
    --name hosts \
    --version #{VERSION} \
    --vendor "Matt Ho" \
    -s dir \
    -t deb \
    -C #{DIST} \
    usr 
EOF
end

def prepare_package 
  exec "mkdir -p #{DIST}/usr/local/bin"
  exec "cp hosts/hosts #{DIST}/usr/local/bin"
end

def exec command
  puts command
  raise "unable to execute command, #{command}" unless system command
end

