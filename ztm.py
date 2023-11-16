#!/usr/bin/python3
import os
import sys
import time
import asyncio
import tomllib
import argparse

import tqdm
import httpx
RootPath = "/home/zer0-hex/.ztm/"

ConfigPath = os.path.join(RootPath, "config.toml")
VersionInfoPath = os.path.join(RootPath, "cersion.list")
DownPath = os.path.join(RootPath, "downloads")

PROXY = {

    'all://': "http://192.168.198.1:10809",

}


def banner():

    version = "v.0.0.1"

    banner = f"""

███████╗████████╗███╗   ███╗

╚══███╔╝╚══██╔══╝████╗ ████║

  ███╔╝    ██║   ██╔████╔██║

 ███╔╝     ██║   ██║╚██╔╝██║

███████╗   ██║   ██║ ╚═╝ ██║

╚══════╝   ╚═╝   ╚═╝     ╚═╝ 

    Zer0-hex      {version}"""

    print(banner)

def flag():

    banner()

    parser = argparse.ArgumentParser(prog='ztk', description='Zer0-hex Tools Manager')
    parser.add_argument('-d', '--down', action='store_true', default=False, help='重新下载所有软件')

    return parser.parse_args()


def loadConfig(configpath: str) -> dict:
    with open(configpath, "rb") as f:

        data = tomllib.load(f)

    return data


async def download(config: dict):

    client = httpx.AsyncClient(proxies=PROXY)
    async def down(url, name, path):
        try:
            response = await client.get(url=url)
            url = response.headers.get('location')
            response = await client.get(url=url)
        except:
            print(f'[-] fail download: {name}')
            return
        save(name=name, content=response.content, path=path)

    for i in config['Tool']:
        for j, v in enumerate(i['Link']):
            url = v
            name = i['Files'][j]
            path = DownPath + '/' + i['Name']
            await down(url=url, name=name, path=path)


def save(content: bytes, name: str, path: str):
    if not os.path.exists(path=path):
        os.makedirs(path, mode=0o0755)
    filepath = path + '/' + name
    with open(filepath, 'wb') as f:
        f.write(content) 
        print('[+] Save to', filepath)


async def syncVersion(config: dict):

    client = httpx.AsyncClient(proxies=PROXY)

    async def getVersion(url):

        response = await client.get(url)

        try:

            url = response.headers.get('location')

            version = url.split('/')[-1]

        except:

            print(f'[-] fail get version: {url}')

        return version


    for i, v in enumerate(config['Tool']):

        version = await getVersion(url=v['Url'])

        print(f'[+] Get {v["Name"]} Version: {version}')

        v['Version'] = version


def getDownLink(config: dict):

    for i in config['Tool']:

        i['Link'] = []

        for index, j in enumerate(i['Files']):

            if j == 'master':

                i['Files'][index] = i['Name']

                link = i['Url'].replace('releases/latest', f'archive/refs/tags/{i["Version"]}.zip')

                i['Link'].append(link)
            else:

                i['Files'][index] = j.replace('VERSION', i['Version'])

                link = i['Url'].replace('latest', f'download/{i["Version"]}/{i["Files"][index]}')

                i['Link'].append(link)


def action(config: dict):
    for i in config['Tool']:
        tag = i['Tag']
        path = os.path.join(RootPath, tag)
        if not os.path.exists(path=path):
            os.makedirs(path, mode=0o0755)
        dirpath = os.path.join(DownPath,  i['Name'])
        for j in i['Files']:
            filepath = os.path.join(dirpath, j)
            action = i['Action'].replace('DIRPATH', dirpath).replace('TAG', tag).replace("FILEPATH", filepath).replace("FILENAME", j).replace("NAME", i['Name'])
            print('[+]', action)
            os.system(action)

def check():
    pass

def main():
    args = flag()
    config = loadConfig(ConfigPath)
    asyncio.run(syncVersion(config=config))
    if args.down:
        check(config)
    getDownLink(config=config)
    asyncio.run(download(config=config))
    action(config)

main()
