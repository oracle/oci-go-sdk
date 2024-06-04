set -e

# TODO: fix flake8 problems in sdk_regions_updater, python_cli, and team_city_scripts
# flake8Excludes: "./venv,./temp,./input_ocibuild,./output_ocibuild*"
flake8Excludes="./venv,./temp,./input_ocibuild,./output_ocibuild*,./sdk_regions_updater,./python_cli,./team_city_scripts"

# TODO: fix these problems so we don't have to ignore the errors
flake8IgnoredErrors="N806,N802,N803,N817,E501,E128,E241,E231,W291,W293"

python -m flake8 --exclude=${flake8Excludes} --ignore=${flake8IgnoredErrors}

pytest tests/