import httplib2
import hmac
import hashlib
import time
import sys
import struct
import json

root = "https://api.challenge.hennge.com/challenges/003"
content_type = "application/json"
userid = "me@tuhidulhossain.com"
secret_suffix = "HENNGECHALLENGE003"
shared_secret = userid+secret_suffix

timestep = 30
T0 = 0

def HOTP(K, C, digits=10):
    """HTOP:
    K is the shared key
    C is the counter value
    digits control the response length
    """
    K_bytes = str.encode(K)
    C_bytes = struct.pack(">Q", C)
    hmac_sha512 = hmac.new(key = K_bytes, msg=C_bytes, digestmod=hashlib.sha512).hexdigest()
    return Truncate(hmac_sha512)[-digits:]

def Truncate(hmac_sha512):
    """truncate sha512 value"""
    offset = int(hmac_sha512[-1], 16)
    binary = int(hmac_sha512[(offset *2):((offset*2)+8)], 16) & 0x7FFFFFFF
    return str(binary)

def TOTP(K, digits=10, timeref = 0, timestep = 30):
    """TOTP, time-based variant of HOTP
    digits control the response length
    the C in HOTP is replaced by ( (currentTime - timeref) / timestep )
    """
    C = int ( time.time() - timeref ) // timestep
    return HOTP(K, C, digits = digits)

data = {  "github_url": "https://gist.github.com/encryptedtouhid/7ac21cd375fc1a4eeeaf53d0e5798f3c",  "contact_email": "me@tuhidulhossain.com" }

passwd = TOTP(shared_secret, 10, T0, timestep).zfill(10) 
print("Password :" + passwd)
# h = httplib2.Http()
# h.add_credentials( userid, passwd )
# header = {"content-type": "application/json"}
# resp, content = h.request(root, "POST", headers = header, body = json.dumps(data))
# print(resp)
# print(content)
