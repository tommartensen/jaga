import time

SECONDS_PER_HOUR = 60 * 60
SECONDS_PER_MINUTE = 60
TIME_FORMAT_HRS_MIN_SEC = "%H:%M:%S"
TIME_FORMAT_MIN_SEC = "%M:%S"


def build_href(href):
    return href.replace("strava://", "https://www.strava.com/")


def _remove_leading_zeroes(string):
    return string.lstrip("0")


def _contains_hours(time_string):
    return len(time_string.split(":")) == 3  # H:M:S


class ElapsedTimeHelper:
    def __init__(self, time_obj):
        self.container = time_obj

    def as_string(self):
        time_string = time.strftime(TIME_FORMAT_MIN_SEC, self.container)
        formatted_time_string = _remove_leading_zeroes(time_string)
        return formatted_time_string

    @classmethod
    def from_seconds(cls, seconds):
        time_obj = time.gmtime(seconds)
        return cls(time_obj)

    @classmethod
    def from_string(cls, time_string):
        if _contains_hours(time_string):
            time_obj = time.strptime(time_string, TIME_FORMAT_HRS_MIN_SEC)
        else:
            time_obj = time.strptime(time_string, TIME_FORMAT_MIN_SEC)
        return cls(time_obj)

    @classmethod
    def build_pace_string(cls, seconds):
        time_struct = cls.from_seconds(seconds)
        return time_struct.as_string()

    @property
    def elapsed_seconds(self):
        return (
            self.container.tm_hour * SECONDS_PER_HOUR
            + self.container.tm_min * SECONDS_PER_MINUTE
            + self.container.tm_sec
        )

    @property
    def elapsed_hours(self):
        return self.elapsed_seconds / SECONDS_PER_HOUR
