import colorama
import functools
import os
import signal
import time
from typing import Tuple

from webtree import Page, Tree


class Crawler:
    def __init__(self, root_url: str, level: int, live_mode: bool, export_file: str, regex_map: dict, excluded_status: list, included_urls: list) -> None:
        self.root_url = root_url
        self.level = level
        self.live_mode = live_mode
        self.export_file = export_file
        self.regex_map = regex_map
        self.excluded_status = excluded_status
        self.included_urls = included_urls
        self._cache = Cache()

    @functools.cached_property
    def cache(self) -> Cache:
        return self._cache

    def crawl(self) -> None:
        root = Tree(Page(url=self.root_url))
        interrupt_chan = signal.signal(signal.SIGINT, signal.SIG_IGN)
        try:
            if self.live_mode:
                self._crawl_live(root)
            else:
                self._crawl_block(root)
        except KeyboardInterrupt:
            pass
        finally:
            signal.signal(signal.SIGINT, interrupt_chan)

    def _crawl_live(self, root: Tree) -> None:
        spinner = colorama.Fore.YELLOW + "incursion ... " + colorama.Style.RESET_ALL
        spinner = colorama.AnsiColor.render(spinner)
        for _ in range(self.level):
            print(spinner, end="")
            time.sleep(0.1)
        for node in root.traverse():
            node.page.fetch()
            node.page.print_page_live()
            self._cache.add_visited(node.page.url)
            for link in node.page.links:
                if not self._cache.is_visited(link):
                    node.add_child(Page(url=link))

    def _crawl_block(self, root: Tree) -> None:
        for i in range(self.level, 0, -1):
            for node in root.traverse(i):
                node.page.fetch()
                self._cache.add_visited(node.page.url)
                for link in node.page.links:
                    if not self._cache.is_visited(link):
                        child = Page(url=link)
                        self._crawl_block(child)

