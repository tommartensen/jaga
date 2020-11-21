from jaga.api import SegmentsApi
from jaga.utils import ElapsedTimeHelper


class Segment:
    api_response = {}

    def __init__(self, segment_id):
        self.api_response = SegmentsApi().get_by_id(segment_id)
        self.name = self.api_response["name"]
        self.kom_time = ElapsedTimeHelper.from_string(self.api_response["xoms"]["kom"])
        self.length = self.api_response["distance"]
        self.average_gradient = self.api_response["average_grade"]
        self.href = self._build_href(self.api_response["xoms"]["destination"]["href"])

    @property
    def length_in_km(self):
        return self.length / 1000

    @property
    def kom_pace(self):
        pace_in_seconds = self.kom_time.elapsed_seconds / self.length_in_km
        return ElapsedTimeHelper.build_pace_string(pace_in_seconds)

    @property
    def kom_speed(self):
        return self.length_in_km / self.kom_time.elapsed_hours

    def _build_href(self, href):
        return href.replace("strava://", "https://www.strava.com/")

    def as_string(self):
        return {
            "name": self.name,
            "length": self.length_in_km,
            "kom_time": self.kom_time.as_string(),
            "kom_pace": self.kom_pace,
            "kom_speed": self.kom_speed,
            "average_gradient": self.average_gradient,
            "link": self.href,
        }
