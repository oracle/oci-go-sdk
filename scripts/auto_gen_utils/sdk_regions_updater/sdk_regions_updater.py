import argparse
import configparser
import logging
import os
import sys
from .dotnet_sdk_region_updater import DotNetSDKRegionsUpdater
from .go_sdk_region_updater import GoSDKRegionsUpdater
from .java_sdk_region_updater import JavaSDKRegionsUpdater
from .python_sdk_region_updater import PythonSDKRegionsUpdater
from .ruby_sdk_region_updater import RubySDKRegionsUpdater
from .typescript_sdk_region_updater import TypescriptSDKRegionsUpdater
from itertools import chain
from .region_updater_utils import get_new_regions_info_for_branch

# Add the root of the package, two directories up, to the sys.path
dir_path = os.path.dirname(os.path.realpath(__file__))
sys.path.append(os.path.join(dir_path, '..'))
import config # noqa: ignore=F402


class SdkRegionsUpdater:
    def __init__(self, sdk=None, branch='preview'):
        self._sdk_updater_configuration(sdk, branch)
    
    def _sdk_updater_configuration(self, sdk, branch):
        if sdk is None:
            args_parser = argparse.ArgumentParser()
            args_parser.add_argument('--SDK', 
                                    required=True,
                                    type=str, 
                                    choices=config.SDKS_SUPPORTING_REGION_UPDATE + ['All'],
                                    dest="sdk")
            args_parser.add_argument('--Branch',
                                    default='preview',
                                    type=str,
                                    dest='branch',
                                    help='The branch to update (default = preview).')
            args = args_parser.parse_args()
            self.sdk = args.sdk
            self.branch = args.branch
        else:
            self.sdk = sdk
            self.branch = branch

        config_parser = configparser.ConfigParser()
        script_dir = os.path.dirname(os.path.realpath(__file__))
        with open(os.path.join(script_dir, "config")) as lines:
            lines = chain(("[default]",), lines)
            config_parser.read_file(lines)
        self.config = config_parser['default']
        self.logger = logging.getLogger("OCI-SDK-Regions-Updater")
        logging.basicConfig(level=logging.DEBUG)

    def _get_config_value(self, key):
        val = os.environ.get(key)
        if val is None:
            self.logger.log(logging.INFO, 'Unable to get value for {} from env variable. Will try to get value from config file.'.format(key))
            val = self.config.get(key)
        if val is None or val == '':
            raise ValueError('Unable to get value for {} from env variable or config file.'.format(key))
        else:
            return val

    def _validate_path_exists(self, path):
        if os.path.exists(path):
            self.logger.log(logging.INFO, '{} is a valid path'.format(path))
        else:
            raise ValueError('{} is not a valid path'.format(path))

    def _update_sdk(self, sdk, new_regions):
        self.logger.log(logging.INFO, 'Updating {}'.format(sdk))
        sdk_path = self._get_config_value('{}_path'.format(sdk))
        self._validate_path_exists(sdk_path)
        if sdk == config.DOTNET_SDK_NAME:
            sdk_updater = DotNetSDKRegionsUpdater(sdk_path)
        elif sdk == config.JAVA_SDK_NAME:
            sdk_updater = JavaSDKRegionsUpdater(sdk_path)
        elif sdk == config.TYPESCRIPT_SDK_NAME:
            sdk_updater = TypescriptSDKRegionsUpdater(sdk_path)
        elif sdk == config.PYTHON_SDK_NAME:
            sdk_updater = PythonSDKRegionsUpdater(sdk_path)
        elif sdk == config.GO_SDK_NAME:
            sdk_updater = GoSDKRegionsUpdater(sdk_path)
        elif sdk == config.RUBY_SDK_NAME:
            sdk_updater = RubySDKRegionsUpdater(sdk_path)
        else:
            raise ValueError('{} is not supported.'.format(sdk))
        sdk_updater.update_regions(new_regions)

    def update(self, new_regions=None):
        if not new_regions:
            new_regions = get_new_regions_info_for_branch(self.branch)
        if len(new_regions) == 0:
            self.logger.log(logging.INFO, 'No new regions to update.')
        elif self.sdk.lower() == 'all':
            for supported_sdk in config.SDKS_SUPPORTING_REGION_UPDATE:
                self._update_sdk(supported_sdk, new_regions)
        else:
            self._update_sdk(self.sdk, new_regions)
        return len(new_regions)


if __name__ == "__main__":
    updater = SdkRegionsUpdater()
    num_regions_updated = updater.update()
    print('{} regions updated.'.format(num_regions_updated))
