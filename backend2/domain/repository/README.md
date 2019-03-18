# repositoryパッケージ

[DDD](http://domainlanguage.com/wp-content/uploads/2016/05/DDD_Reference_2015-03.pdf)における「Repositories」に相当する。

集約の永続化管理を担当

エンティティや値オブジェクトから構成される集約の格納と取得を担当

通常、集約とリポジトリの関係は一対一になる

いわゆるDAO(DataAccessObject)に似ているが、DAOがデータ指向であるのに対し、Repositoryはオブジェクト指向でのアプローチ。

インタフェースのみ定義する。
