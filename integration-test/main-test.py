import requests
import unittest


if __name__ == '__main__':
    unittest.main()


class Test(unittest.TestCase):
    def test_health(self):
        response = requests.get("http://localhost:8090/health")
        self.assertEqual(200, response.status_code)