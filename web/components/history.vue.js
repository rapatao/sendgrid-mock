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
      }
    },
    open(id, format) {
      window.open('/messages/' + id + '?' + new URLSearchParams({
        "format": format
      }))
    },
    confirmDelete(id) {
      if (confirm("Are you sure you want to delete this message?")) {
        this.deleteFunc(id)
      }
    },
    hasContent(content) {
      return content && content !== ""
    },
    messageCount(state) {
      let messagesStart = state.page * state.maxRows

      let firstMessage = 1 + messagesStart
      if (state.messages.length === 0) {
        firstMessage = 0
      }

      let lastMessage = (state.page + 1) * state.maxRows
      if (state.messages.length < state.maxRows) {
        lastMessage = messagesStart + state.messages.length
      }

      return `${firstMessage} to ${lastMessage} of ${state.total} message(s).`
    }
  }
  ,
  template: `
    <div class="box mt-2">
      <progress v-if="state.loading" class="progress is-small is-link mb-2" max="100">Loading</progress>
      <h2 class="title is-5 mb-3">Message History</h2>
      <div class="table-container">
        <table class="table is-fullwidth is-hoverable is-striped table-history">
          <thead>
          <tr>
            <th style="width: 15%">Date</th>
            <th style="width: 20%">From</th>
            <th style="width: 20%">To</th>
            <th>Subject</th>
            <th class="has-text-centered" style="width: 15%">Actions</th>
          </tr>
          </thead>

          <tbody>
          <tr v-for="message in state.messages" :key="message.event_id">
            <td class="is-vcentered">
              {{ new Date(message.received_at).toLocaleString() }}
            </td>

            <td class="is-vcentered">
              <p class="is-marginless" v-if="message.from.name"><strong>{{ message.from.name }}</strong> &lt;{{ message.from.address }}&gt;</p>
              <p class="is-marginless" v-else><strong>{{ message.from.address }}</strong></p>
            </td>

            <td class="is-vcentered">
              <p class="is-marginless" v-if="message.to.name"><strong>{{ message.to.name }}</strong> &lt;{{ message.to.address }}&gt;</p>
              <p class="is-marginless" v-else><strong>{{ message.to.address }}</strong></p>
            </td>

            <td class="is-vcentered">{{ message.subject }}</td>

            <td class="is-vcentered has-text-centered">
              <div class="buttons is-centered">
                <button class="button is-small is-info is-light" 
                        @click="open(message.event_id, 'html')" 
                        v-show="hasContent(message.content.html)"
                        title="View HTML version">
                  <span class="icon"><i class="fas fa-code"></i></span>
                </button>

                <button class="button is-small is-info is-light" 
                        @click="open(message.event_id, 'text')" 
                        v-show="hasContent(message.content.text)"
                        title="View Text version">
                  <span class="icon"><i class="fas fa-file-alt"></i></span>
                </button>

                <button class="button is-small is-danger is-light" 
                        @click="confirmDelete(message.event_id)"
                        title="Delete message">
                  <span class="icon"><i class="fas fa-trash"></i></span>
                </button>
              </div>
            </td>
          </tr>
          <tr v-if="state.messages.length === 0">
            <td colspan="5" class="has-text-centered py-6">
              <span class="icon is-large has-text-grey-light"><i class="fas fa-3x fa-inbox"></i></span>
              <p class="is-size-5 has-text-grey-light mt-4">No messages found</p>
            </td>
          </tr>
          </tbody>
        </table>
      </div>

      <div class="level mt-4">
        <div class="level-left">
          <div class="level-item">
            <p class="is-size-7 has-text-grey">
              {{ messageCount(state) }}
            </p>
          </div>
        </div>
        <div class="level-right">
          <div class="level-item">
            <div class="field has-addons">
              <p class="control">
                <button class="button is-small" :disabled="!hasPrevious(state)" @click="previous">
                  <span class="icon is-small"><i class="fas fa-chevron-left"></i></span>
                  <span>Previous</span>
                </button>
              </p>
              <p class="control">
                <button class="button is-small" :disabled="!hasNext(state)" @click="next">
                  <span>Next</span>
                  <span class="icon is-small"><i class="fas fa-chevron-right"></i></span>
                </button>
              </p>
            </div>
          </div>
        </div>
      </div>
    </div>
  `
}
