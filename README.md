![Thumbnail](https://lh5.googleusercontent.com/proxy/r5D7LX7gbvXfuJU1SFAfCM1SerPt0KcBvR_R0qpXO_fsa39nwCKhyGE0UQbFP99XpSMRuPWrckLRnkoU747FW6EHY1_Gqf1xzhXYhJnIqIHizuhbBX3fh0sgdxbpIwJrDtC9g-uELzM-xYNfiw=s0-d)



# INSTAGRAM CRAWLER

[![LICENSE](https://img.shields.io/badge/license-Apache%202-blue)](https://github.com/joshua-dev/instagram-crawler/blob/master/LICENSE)
[![go version](https://img.shields.io/badge/go-1.14-00ADD8)](https://go.dev)



## :white_check_mark: TODO

* ### Model

  - [x] TopSearch: 검색어를 search 칸에 입력했을 때 나오는 태그 목록과 게시물 갯수를 반환
  - [ ] Crawl: query(검색어)로 태그 검색 시 나오는 각각의 게시물 (https://instagram.com/explore/tags/{query} 의 결과) 의 URL, 커버 사진 URL, 태그 정보, 좋아요 갯수를 게시물 구조체 자료구조에 담아 반환
    - [ ] CrawlNumCoreSpriteHashtag: 태그 검색 url의 게시물 갯수를 반환 (접근 클래스 이름 = ".g47SY")
      - [ ] CoreSpriteHashtag: 각 게시물 url마다 접근, 커버 사진 URL, 태그 정보, 좋아요 갯수를 구조체에 담아 반환 (접근 클래스 이름 = ".v1Nh3")
        - [ ] CoreSpriteHashtagSrc: 커버 사진 URL을 반환 (접근 클래스 이름 = "._9AhH0")
        - [ ] CoreSpriteHashtagTagsInfo: 태그 정보를 반환 (접근 클래스 이름 = ".xil3i")
        - [ ] CoreSpriteHashtagLike: 좋아요 갯수를 반환 (접근 클래스 이름 = ".sqdOP")

* ### View
  
  - [ ] API 서버 구축

* ### Controller
  
  - [ ] 게시물 구조체에 대한 String, MarshalJSON 구현


- [ ] 최적화: 병렬 처리

- [ ] API 명세서 작성

---

### [LICENSE](https://github.com/joshua-dev/instagram-crawler/blob/master/LICENSE)
