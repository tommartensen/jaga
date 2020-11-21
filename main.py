import sys
from pprint import pprint

from jaga.models import Segment

if __name__ == "__main__":
    segment_id = sys.argv[1]
    pprint(Segment(segment_id).as_string())
