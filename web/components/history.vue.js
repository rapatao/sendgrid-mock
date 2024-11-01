export default {
  props: {
    state: Object,
    deleteEventFunc: Function,
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
              {{ message.received_at }}
            </td>

            <td>
              <strong>{{ message.from_name }}</strong>
              <br/>
              <small>{{ message.from_address }}</small>
            </td>

            <td>
              <strong>{{ message.to_name }}</strong>
              <br/>
              <small>{{ message.to_address }}</small>
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
          <a href="#" class="pagination-previous is-disabled">Previous</a>
          <a href="#" class="pagination-next is-disabled">Next page</a>
        </nav>
      </div>
    </section>
  `
}
