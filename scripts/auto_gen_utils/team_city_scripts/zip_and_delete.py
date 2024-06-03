from __future__ import print_function

import argparse
import os
import zipfile
import shutil

parser = argparse.ArgumentParser(description='Script to zip and delete files')
parser.add_argument('-d', '--directory', type=str, required=True, help='Directory path to be zipped and deleted')
parser.add_argument('--verbose', default=False, action='store_true', help='Verbose logging')
parser.add_argument('--dry-run', default=False, action='store_true', help='Dry-run, do not delete files')

args = parser.parse_args()
dir_path = args.directory
is_verbose = args.verbose
dry_run = args.dry_run

def printv(s):
    if is_verbose:
        print(s)

def zip_and_delete_files(foldername):
    source_dir = foldername
    zip_dir = foldername + '.zip'
    zipobj = zipfile.ZipFile(zip_dir, 'w', zipfile.ZIP_DEFLATED)
    rootlen = len(source_dir) + 1
    for base, dirs, files in os.walk(source_dir):
        for file in files:
            file_name = os.path.join(base, file)
            printv("Adding file to zip: {}".format(file_name))
            zipobj.write(file_name, file_name[rootlen:])
            if not dry_run:
                printv("Removing Zipped file: {}".format(file_name))
                os.remove(file_name)
    if not dry_run:
        printv("Removing Dir: {}".format(source_dir))
        shutil.rmtree(source_dir)

printv("Zipping directory: {}".format(dir_path))
zip_and_delete_files(dir_path)