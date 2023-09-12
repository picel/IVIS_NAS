# IVIS NAS Internal Web Server

## 개요
- IVIS Lab 내부망에서 동작하는 웹서버.
- NAS의 사용자 추가 및 메일 사용자 추가를 위한 웹페이지 제공.
- [Helper 서버](https://github.com/picel/IVIS_NAS_Helper)에서의 사용자 인증 API 제공.

## 전체 시스템 개략도
![ivis_eco](https://github.com/picel/IVIS_NAS/assets/30901178/224966f4-e613-4da7-90ec-5175e1b2a07e)

## 개발환경
- Go
- MXRoute Mail API (DirectAdmin API)
- SMB
- AFP
- NGINX (Reverse Proxy)

## 라우트 정보
- GET /
    - NAS 사용자 추가 페이지
- POST /process
    - NAS 사용자 추가 처리 페이지
    - UNIX 계정 생성
    - SMB 계정 생성
    - MXRoute Mail 사용자 자동 추가
- POST /loginCheck
    - Helper 서버에서 사용자 인증 요청용 API
    - Helper 서버에서 사용자 인증 요청 시, UNIX 계정의 존재 여부를 확인하여 인증 결과를 반환.
    - /etc/shadow 파일을 읽어와서 인증 처리.

## 실행 화면
- NAS 사용자 추가 페이지
![NAS 사용자 추가 페이지](https://user-images.githubusercontent.com/30901178/222890273-194bfde2-8ca2-4c21-b972-72a6d42de615.png)
- NAS 사용자 등록 결과 페이지
![NAS 사용자 등록 결과 페이지](https://user-images.githubusercontent.com/30901178/222890305-ff4eb233-0a31-48ec-b486-c6921b474bb3.png)

## 주의사항
- key 파일은 git에 올리지 않음.
    - key 파일은 /etc/goweb/keys 디렉토리에 위치.
    - MXRouteURL, MXRouteID, MXRoutePW 정의
