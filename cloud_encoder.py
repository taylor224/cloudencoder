#!/usr/bin/python3
"""
OpenEncoder 를 ffmpeg binary 처럼 사용 가능하게 해주는 Python Script

alias ffmpeg = "python3 cloud_encoder.py"
하셔서 사용하시면 됩니다.

단 parameter 는 input 파일 path와 output 파일 path만 받아서 OpenEncoder 에서 미리 처리된 방식으로 인코딩 하니 주의하시기 바랍니다.

by. Taylor
"""
import argparse
import requests
import time
import os


TOKEN = ''
DOWNLOAD_BASE_URL = 'http://YOUR_SERVER_WEB_SERVER_URL'
BASE_URL = 'http://ENCODING_SERVER_API_URL:8080'
ID = 'ENCODING_USER_ID'
PASSWORD = 'ENCODING_USER_PASSWORD'
DEFAULT_ENCODING_PRESET = 'mastodon_h264_encode'
SMALL_ENCODING_PRESET = 'mastodon_h264_small'

parser = argparse.ArgumentParser(
                    prog = 'Cloud Encoder',
                    description = 'FFmpeg to Cloud Encode',
                    epilog = './cloud_encoder -i input.mp4 output.mp4')

parser.add_argument('-i', '--input')
parser.add_argument('-p', '--preset')
parser.add_argument('-y', '--yes')


def send_encode(file_path, preset):
  global TOKEN
  try:
    login_data = {
      'username': ID,
      'password': PASSWORD
    }
    r = requests.post('%s/api/login' % BASE_URL, json=login_data)
    if r.status_code != 200:
      print(r.content)
      return False

    jwt_data = r.json()
    TOKEN = jwt_data['token']
    headers = {
      'Authorization': 'Bearer %s' % jwt_data['token']
    }
    with open(file_path, 'rb') as f:
      files = {
        'file': f
      }
      r = requests.post('%s/api/upload' % BASE_URL, headers=headers, files=files)
      if r.status_code != 200:
        print(r.content)
        return False

      data = r.json()

      if preset == 'small':
        request_preset = SMALL_ENCODING_PRESET
      else:
        request_preset = DEFAULT_ENCODING_PRESET

      job_data = {
        'preset': request_preset,
        'source': '/tmp/uploads/%s' % data['file_name'],
      }

      r = requests.post('%s/api/jobs' % BASE_URL, headers=headers, json=job_data)
      if r.status_code != 201:
        print(r.content)
        return False
      return r.json()
  except Exception as e:
    print(e)
    return False


def check_status(data):
  global TOKEN
  headers = {
    'Authorization': 'Bearer %s' % TOKEN
  }
  job_id = data['job']['id']
  guid = data['job']['guid']

  for i in range(100):
    r = requests.get('%s/api/jobs/%d/status' % (BASE_URL, job_id), headers=headers)
    if r.status_code != 200:
      print(r.content)
      time.sleep(1)
      continue

    job_status = r.json()['job_status']
    if job_status == 'error':
      return False
    elif job_status == 'completed':
      return True
    time.sleep(1)

  return False


def download_file(data, output):
  r = requests.get('%s/results/%s/mastodon_h264_encode.mp4' % (DOWNLOAD_BASE_URL, data['job']['guid']), stream=True)
  if r.status_code != 200:
    print(r.content)
    return False

  with open(output , 'wb') as f:
    for chunk in r.iter_content(1024):
      f.write(chunk)
  r.close()
  return True


if __name__ == '__main__':
  DOWNLOAD_BASE_URL = os.getenv('ENCODER_DOWNLOAD_BASE_URL', DOWNLOAD_BASE_URL)
  BASE_URL = os.getenv('ENCODER_BASE_URL', BASE_URL)

  args, unknown = parser.parse_known_args()
  data = send_encode(args.input, args.preset)
  if not data:
    exit(1)
  if check_status(data):
    if download_file(data, args.yes):
      exit(0)
    else:
      exit(1)
  else:
    exit(1)
