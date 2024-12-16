<template>
    <div class="scenic-spots-container">
        <h1>景区介绍</h1>
        <div v-for="item in this.$store.state.Scenicspots" :key="item.scenicspotid" class="scenic-spot">
            <div class="scenic-spot-header" @click="toggleDetails(item.scenicspotid)">
                <h2>{{ item.scenicspotname }}</h2>
                <p class="scenic-spot-id">ID: {{ item.scenicspotid }}</p>
            </div>
            <p class="scenic-spot-details" :class="{ 'expanded': expandedSpots[item.scenicspotid] }">
                {{ item.scenicspotdetails }}
            </p>
        </div>
    </div>
</template>

<script>
export default {
    data() {
        return {
            expandedSpots: {}
        };
    },
    methods: {
        toggleDetails(spotId) {
            this.$set(this.expandedSpots, spotId, !this.expandedSpots[spotId]);
        }
    }
};
</script>

<style scoped>
.scenic-spots-container {
    max-width: 800px;
    margin: 0 auto;
    padding: 20px;
}

.scenic-spot {
    background-color: #f9f9f9;
    border-radius: 8px;
    margin-bottom: 20px;
    overflow: hidden;
    box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.scenic-spot-header {
    background-color: #4a90e2;
    color: white;
    padding: 15px;
    cursor: pointer;
    display: flex;
    justify-content: space-between;
    align-items: center;
}

.scenic-spot-header h2 {
    margin: 0;
    font-size: 1.2em;
}

.scenic-spot-id {
    font-size: 0.9em;
    opacity: 0.8;
}

.scenic-spot-details {
    padding: 15px;
    margin: 0;
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
    transition: max-height 0.3s ease-out, padding 0.3s ease-out;
    max-height: 1.5em;
}

.scenic-spot-details.expanded {
    white-space: normal;
    overflow: visible;
    text-overflow: clip;
    max-height: 1000px;
    /* 一个足够大的值来容纳完整内容 */
}
</style>