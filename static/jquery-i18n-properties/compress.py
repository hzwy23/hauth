#!/usr/bin/python

import httplib, urllib, sys

# Define the parameters for the POST request and encode them in
# a URL-safe format.
f = open('jquery.i18n.properties.js', 'r')

params = urllib.urlencode([
    #('js_code', sys.argv[1]),
    ('js_code', f.read()),
    ('compilation_level', 'SIMPLE_OPTIMIZATIONS'),
    ('output_format', 'text'),
    ('output_info', 'compiled_code'),
  ])

# Always use the following value for the Content-type header.
headers = { "Content-type": "application/x-www-form-urlencoded" }
conn = httplib.HTTPConnection('closure-compiler.appspot.com')
conn.request('POST', '/compile', params, headers)
response = conn.getresponse()
data = response.read()

#print data
f2 = open('jquery.i18n.properties.min.js', 'w')
f2.write(data)

conn.close
