from time import sleep
import sys
import urllib.request

if len(sys.argv) < 1:
    print("Please provide a URL to healthcheck ;)")
    exit(1)
url = sys.argv[1]

print("SPAMMING:", url)
while 1 > 0:
    try:
        contents = urllib.request.urlopen(str(url))
        print(contents.peek(20))
        sleep(1)
    except:
        print("OUTAGE")
        sleep(1)
