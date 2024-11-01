export default {
  props: {
    state: Object,
    deleteEventFunc: Function,
    filterFunc: Function,
  },
  methods: {
    previous() {
      this.state.page -= 1
      this.filterFunc()
    },
    next() {
      this.state.page += 1
      this.filterFunc()
    },
  },
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
            <span class="icon">
              <span class="has-text-success">
                <i class="fas fa-lg fa-file-code"></i>
              </span>
            </span>
              <span class="icon">
                <span class="has-text-success">
                  <i class="fas fa-lg fa-file-alt"></i>
                </span>
            </span>
              <span class="icon" @click="deleteEventFunc(message.event_id)">
              <span class="has-text-danger">
                <i class="fas fa-lg fa-trash"></i>
              </span>
            </span>
            </td>
          </tr>
          </tbody>
        </table>

        <nav class="pagination" role="navigation" aria-label="pagination">
          <a href="#" class="pagination-previous" :class="{ 'is-disabled': state.page <= 0}" @click="previous()">Previous</a>
          <a href="#" class="pagination-next"
             :class="{ 'is-disabled': (state.total - ((1 + state.page) * state.maxRows)) < 1 }"
             @click="next()">Next page</a>
        </nav>
      </div>
    </section>
  `
}
