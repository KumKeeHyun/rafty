# rafty
합의 알고리즘 `Raft` 구현해보기

`Diego Ongaro`의 `In Search of an Understandable Consensus Algorithm` 논문에 충실하게 구현

## 목표
- 클러스터 구성 변경은 제외 (다음 버전에서 구현)
- 로그 복제, 스냅샷 구현
    - 로그 엔트리들은 메모리에 저장
    - 스냅샷은 파일에 저장. 파일명과 다른 메타데이터를 따로 메모리에 저장
- StateMachine은 인터페이스로 빼놓고 메모리 Key-Value Store 구현
    - 스탭샷 요구사항 : 스냅샷을 만드는 동안에도 업데이트 가능
    - copy-on-write 써야함. 일단 임시로 immutable한 구조체로 구현 
- 클라이언트 요청 구현
    - 클라이언트 코드 라이브러리 형식으로 구현
    - 리더찾기는 클러스터에 랜덤으로 찔러보고 리더가 아닌 서버가 받으면 리더의 주소를 반환해주는 형식으로

## 군바
