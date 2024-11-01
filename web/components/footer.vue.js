export default {
  props: {
    state: Object,
  },
  template: `
    <footer class="footer">
      <div class="content has-text-centered">
        <p><strong>Showing {{ state.messages.length }} of {{ state.total }}</strong> message(s).</p>
      </div>
    </footer>
  `
}
