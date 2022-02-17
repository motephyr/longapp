// 須指定pageNumber = 0才能啟用分頁
// 指定model為即時更新之變數
// setModelVariable為另外的更新方法
// variables為相關的變數內容
let mixin = {
    data() {
        return {
            pageNumber: this.$options.queryModels.pageNumber,
            total: 0,
            pageCount: 0,
            pageList: [],
            setModelFunc: this.$options.queryModels.setModelFunc || function (x) { return x },
            model: this.$options.queryModels.model,
            variables: this.$options.queryModels.variables,
        };
    },
    async created() {
        this.unwatch = this.$watch(this.$options.queryModels.variables, this.$options.queryModels.watch)
    },
    destroyed() {
        this.unwatch()
    },
    methods: {
        async handleResults(result) {
            result.data = result.data.map((x) => {
                return this.setModelFunc(x);
            });
            this[this.model] = result.data;

            if (this.pageNumber != null) {
                this.pageCount = result.lastPage;
                // this.pageNumber = data.candidatesPagination.page
                this.total = result.total;
                this.getPageList()
            }
        },
        async handleResult(result, id) {
            result.data = result.data.map((x) => {
                return this.setModelFunc(x);
            });
            if (this[this.model].find((x) => x.id === id)) {
                this[this.model] = [...this[this.model].map((x) => {
                    return x.id === id ? result.data[0] : x;
                })];
            } else {
                this[this.model] = result.data.concat(this[this.model]);
            }
        },

        getPageList() {
            let list = [];
            let page = this.pageNumber + 1,
                start,
                end;
            if (page <= 5) {
                start = 1;
                end = 10;
            } else if (page > this.pageCount - 5) {
                start = this.pageCount - 9;
                end = this.pageCount;
            } else {
                start = page - 4;
                end = page + 5;
            }

            start = start < 1 ? 1 : start;
            end = end > this.pageCount ? this.pageCount : end;

            for (let i = start; i <= end; i++) {
                list.push(i);
            }
            this.pageList = list;
        },
    }
};

export default mixin;


