# -*- coding: utf-8 -*-
from struct import unpack
from os import urandom
import string


def generate_short_id():
    """
    short id create, only has 11-character
    """
    num = unpack(str("<Q"), urandom(8))[0]
    if num <= 0:
        result = "0"
    else:
        alphabet = string.digits + string.ascii_uppercase + string.ascii_lowercase
        key = []
        while num > 0:
            num, rem = divmod(num, 62)
            key.append(alphabet[rem])
        result = "".join(reversed(key))
    return result
