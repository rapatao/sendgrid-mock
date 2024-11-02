export default {
  props: {
    state: Object,
    deleteAllFunc: Function,
  },
  methods: {
    footer(state) {
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
    },
  },
  template: `
    <footer class="footer">
      <div class="content has-text-centered">
        <p>{{ footer(state) }}</p>
      </div>
      <div class="content has-text-right">
        <a @click="deleteAllFunc">delete all</a>
      </div>
    </footer>
  `
}
