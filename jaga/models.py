from jaga.strava_proxy import SegmentsProxy
from jaga.utils import ElapsedTimeHelper, build_href


class Segment:
    def __init__(self, segment_id):
        api_response = SegmentsProxy().get_by_id(segment_id)
        self.name = api_response["name"]
        self.kom_time = ElapsedTimeHelper.from_string(api_response["xoms"]["kom"])
        self.length = api_response["distance"]
        self.average_gradient = api_response["average_grade"]
        self.href = build_href(api_response["xoms"]["destination"]["href"])

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

    def __repr__(self):
        return f"""
Name: {self.name}
Length: {self.length_in_km:.3f}km
KOM Time: {self.kom_time}
KOM Pace: {self.kom_pace} min/km
KOM Speed: {self.kom_speed:.2f} km/h
Avg Gradient: {self.average_gradient}%
Link: {self.href}
"""
