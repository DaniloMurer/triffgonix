<script setup>

const messages = ref()

onMounted(() => {
  const socket = connectToSocket('201');
  socket.onopen = () => {
    const handshake = {
      type: "handshake",
      data: {
        content: "data"
      }
    };
    const throw1 = {
      type: "throw",
      data: {
        content: "data"
      }
    };
    socket.send(JSON.stringify(handshake));
    console.log("conneted to websocket");
    socket.send(JSON.stringify(throw1));
  }
  socket.onmessage = (evt) => {
    console.log(evt);
    const data = JSON.parse(evt.data);
    messages.value = data.currentPlayer.score;
  }
})
</script>
<template>
  <h1 class="text-center">Hello world</h1>
  <span>{{ messages }}</span>
</template>
