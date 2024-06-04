#!/bin/bash

# ./change_dexreq_to_release_approved.sh " DEXREQ-1803, DEXREQ-1803 "
set -e
set -x
LIST_OF_GO_TICKETS=$1

LIST_OF_GO_TICKETS_ARG=""
if [ ! -z "$LIST_OF_GO_TICKETS" ]; then
  # Remove spaces from LIST_OF_GO_TICKETS. DEXREQ tickets should be only comma seperated
  LIST_OF_GO_TICKETS="$(echo -e "${LIST_OF_GO_TICKETS}" | tr -d '[:space:]')"
  LIST_OF_GO_TICKETS_ARG="--list-of-go-tickets "$LIST_OF_GO_TICKETS
  echo "Running ./change_dexreq_to_release_approved $LIST_OF_GO_TICKETS_ARG"
  python ./change_dexreq_to_release_approved.py $DRY_RUN_ARG $LIST_OF_GO_TICKETS_ARG
fi
