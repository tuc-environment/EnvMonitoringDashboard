<template>
    <div>
        <div class="d-flex justify-content-end align-items-center">
            <div class="my-2 mx-2">Count: 199, Showing {{ offset + 1 }} - {{ offset + limit }}</div>
            <nav>
                <ul class="pagination pagination-sm justify-content-end my-2">
                    <li class="page-item" :class="{ disabled: currentPageIdx == 1 }">
                        <a class="page-link" tabindex="-1">Previous</a>
                    </li>
                    <li class="page-item" v-for="(label, idx) in pageLabels" :key="idx"
                        :class="{ disabled: label == currentPageIdx }">
                        <a class="page-link">{{ label }}</a>
                    </li>
                    <li class="page-item" :class="{ disabled: currentPageIdx == pagesCount }">
                        <a class="page-link" tabindex="-1">Next</a>
                    </li>
                </ul>
            </nav>
        </div>

        <table class="table table-bordered">
            <thead class="table-dark">
                <tr>
                    <th scope="col">#</th>
                    <th scope="col">First</th>
                    <th scope="col">Last</th>
                    <th scope="col">Handle</th>
                </tr>
            </thead>
            <tbody>
                <tr>
                    <td>1</td>
                    <td>Test</td>
                    <td>Test</td>
                    <td>Test</td>
                </tr>
            </tbody>
        </table>
    </div>
</template>

<script lang="ts">
export default {
    props: {
        dataUrl: { type: String, required: true, default: 'http://sample-data/api' }
    },
    async created() {
        await this.request()
    },
    data() {
        return {
            count: 100,
            offset: 21,
            limit: 10
        }
    },
    computed: {
        currentPageIdx(): number {
            return Math.ceil(this.offset / this.limit)
        },
        pagesCount(): number {
            return Math.ceil(this.count / this.limit)
        },
        pageLabels(): Array<number> {
            var labels = []
            for (let i = 1; i <= this.pagesCount; i++) {
                labels.push(i)
            }
            return labels
        }
    },
    methods: {
        async request(url: string) {
            console.log(url)
        },
        apiUrl(offset: number, limit: number, order: string) {
            const url = new URL(this.url)
            url.searchParams.set('offset', offset)
            url.searchParams.set('limit', limit)
            url.searchParams.set('order', order)
            return url.toString()
        },
        async clickPrevious() {
            const offset = this.offset - this.limit
            if (offset < 0) offset = 0
            await this.request(this.apiUrl(offset, this.limit, ''))
        },
        async clickNext() {
            const offset = this.offset + this.limit
            await this.request(this.apiUrl(offset, this.limit, ''))
        }
    }
}
</script>
