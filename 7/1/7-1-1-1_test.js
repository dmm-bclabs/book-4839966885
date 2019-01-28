const Contract = artifacts.require('./Contract.sol') // ①


contract('コントラクト名', accounts => { // ②

    describe('複数のテストケースを包括した文言', function () { // ③

        beforeEach(async function () { // ④
            // it関数実行前に毎回実行される処理
        })

        it('テストケースを表した文言', function () {
            // テストコード
        })

        it('テストケースを表した文言', function () {
            // テストコード
        })
    })

})