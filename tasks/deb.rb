#   Copyright 2013 Matt Ho
#
#   Licensed under the Apache License, Version 2.0 (the "License");
#   you may not use this file except in compliance with the License.
#   You may obtain a copy of the License at
#
#     http:#www.apache.org/licenses/LICENSE-2.0
#
#   Unless required by applicable law or agreed to in writing, software
#   distributed under the License is distributed on an "AS IS" BASIS,
#   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
#   See the License for the specific language governing permissions and
#   limitations under the License.

DIST    = "dist"
VERSION = "1.0"

def create_package
  system <<EOF
  # use fpm to package the scripts
  fpm \
    --force \
    --deb-user 0 \
    --deb-group 0 \
    --url http:#github.com/savaki/awsenv/hosts \
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

