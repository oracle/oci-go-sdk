import sys

from functools import partial
from pprint import PrettyPrinter

debug = False
pp = PrettyPrinter(indent=2)


def debug_print(msg, additional='', pretty=False):
    """ Prints a message to stdout.

    :param str msg: the message to display
    :param str additional: additional text to print
    :param boolean pretty: True if the message is to be pretty-printed; else, False (default)
    """
    if not debug:
        return

    if pretty:
        pp.pprint(msg + additional)
    else:
        print(msg + additional)
    sys.stdout.flush()


def get_partial_debug_print_func(indent):
    return partial(debug_print, ' ' * indent)


def print_wrapper(msg, additional=''):
    print(msg + additional)


def get_partial_print_func(indent):
    return partial(print_wrapper, ' ' * indent)
