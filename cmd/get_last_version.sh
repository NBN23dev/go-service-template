#!/bin/sh
if [ "$ENVIRONMENT" = "production" ]; then
  echo "$GITHUB_REF" | rev | cut -d/ -f1 | rev

  exit 0
fi

echo $(date +%s)

exit 0