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
      <div class="columns">
        <div class="column has-text-left">
          <a href="https://github.com/rapatao/sendgrid-mock" target="_blank">
            <span class="icon"><span class="has-text-success"><i class="fab fa-lg fa-github"></i></span></span>
          </a>
        </div>

        <div class="column has-text-right">
          <a @click="deleteAllFunc">Delete all messages</a>
        </div>
      </div>

    </footer>
  `
}
