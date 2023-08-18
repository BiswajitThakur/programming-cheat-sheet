#!/bin/env python3

import os

def remove_file(file_name: str):
    """ Remove File """
    os.remove(file_name)

def remove_directory(dir_name: str):
    """ Remove Empty Directory """
    os.rmdir(dir_name)

