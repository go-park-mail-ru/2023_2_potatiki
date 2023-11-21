#!/bin/bash
set -e

cp -rT /docker-entrypoint-initdb.d/dictionary /usr/share/postgresql/15/tsearch_data
 