export default {
  props: {
    state: Object,
    filterFunc: Function,
    deleteFunc: Function,
  },
  methods: {
    hasPrevious(state) {
      return state.page > 0
    },
    hasNext(state) {
      let startMessages = (1 + state.page) * state.maxRows
      let remainingMessages = state.total - startMessages

      return remainingMessages > 0
    },
    previous() {
      if (this.hasPrevious(this.state)) {
        this.state.page -= 1
        this.filterFunc()
      }
    },
    next() {
      if (this.hasNext(this.state)) {
        this.state.page += 1
        this.filterFunc()

        this.scrollTop()
      }
    },
    open(id, format) {
      window.open('/messages/' + id + '?' + new URLSearchParams({
        "format": format
      }))
    },
    hasContent(content) {
      return content && content !== ""
    }
  }
  ,
  template: `
    <section class="section">
      <div class="container">
        <table class="table is-fullwidth is-hoverable">
          <thead>
          <tr>
            <th>Date</th>
            <th>From</th>
            <th>To</th>
            <th>Subject</th>
            <th>Actions</th>
          </tr>
          </thead>

          <tbody>
          <tr v-for="message in state.messages">
            <td>
              {{ new Date(message.received_at).toLocaleString() }}
            </td>

            <td>
              <strong>{{ message.from.name }}</strong>
              <br/>
              <small>{{ message.from.address }}</small>
            </td>

            <td>
              <strong>{{ message.to.name }}</strong>
              <br/>
              <small>{{ message.to.address }}</small>
            </td>

            <td>{{ message.subject }}</td>

            <td>
              <a class="button" @click="open(message.event_id, 'html')" v-show="hasContent(message.content.html)">
                <span class="icon"><span class="has-text-success"><i class="fas fa-lg fa-file-code"></i></span></span>
              </a>

              <a class="button" @click="open(message.event_id, 'text')" v-show="hasContent(message.content.text)">
                <span class="icon"><span class="has-text-success"><i class="fas fa-lg fa-file-alt"></i></span></span>
              </a>

              <a class="button" @click="deleteFunc(message.event_id)">
                <span class="icon"><span class="has-text-danger"><i class="fas fa-lg fa-trash"></i></span></span>
              </a>
            </td>
          </tr>
          </tbody>
        </table>

        <nav class="pagination" role="navigation" aria-label="pagination">
          <a class="pagination-previous"
             :class="{ 'is-disabled': !hasPrevious(state) }" @click="previous">Previous</a>
          <a class="pagination-next"
             :class="{ 'is-disabled': !hasNext(state) }"
             @click="next">Next page</a>
        </nav>
      </div>
    </section>
  `
}
