import requests

from jaga import config


class StravaProxy:
    @classmethod
    def get(cls, path):
        headers = {"Authorization": f"Bearer {config.ACCESS_TOKEN}"}
        url = f"{config.BASE_URL}/{path}"
        return requests.get(url, headers=headers).json()


class SegmentsProxy:
    def get_by_id(self, segment_id):
        return StravaApi.get(f"segments/{segment_id}")

    def get_in_perimeter(self, center, radius):
        pass

    def get_in_box(self, start_point, end_point):
        pass
