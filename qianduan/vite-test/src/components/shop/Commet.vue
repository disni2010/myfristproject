<template>
    <div class="comment-section">
        <div class="comment-container">
            <h1 class="section-title">用户反馈</h1>

            <el-form :model="comments" :rules="rules" ref="commentFormRef" class="comment-form">
                <el-form-item prop="comment">
                    <el-input v-model="comments.comment" type="textarea" :rows="4" placeholder="分享您对系统的建议和想法..."
                        class="comment-textarea"></el-input>
                </el-form-item>

                <el-form-item prop="rating" class="rating-item">
                    <div class="rating-label">系统满意度评分</div>
                    <el-rate v-model="comments.rating" :max="5" :colors="['#8C8C8C', '#F5A623', '#4CAF50']" show-text
                        :texts="['非常不满', '不满', '一般', '满意', '非常满意']" class="custom-rate"></el-rate>
                </el-form-item>

                <el-form-item>
                    <el-button type="primary" @click="submitComment"
                        :disabled="!comments.comment || comments.rating === 0" class="submit-button">
                        提交反馈
                    </el-button>
                </el-form-item>
            </el-form>

            <el-alert v-if="submissionMessage" :title="submissionMessage"
                :type="isSubmissionSuccess ? 'success' : 'error'" show-icon class="submission-alert"></el-alert>
        </div>
    </div>
</template>

<script>
export default {
    data() {
        return {
            comments: {
                userid: this.$store.state.mainuser.user.userid,
                comment: '',
                rating: 0,
                date: ''
            },
            submissionMessage: '',
            isSubmissionSuccess: false,
            rules: {
                comment: [
                    {
                        required: true,
                        message: '请输入评论内容',
                        trigger: 'blur'
                    },
                    {
                        min: 5,
                        max: 500,
                        message: '评论长度在 5 到 500 个字符之间',
                        trigger: 'blur'
                    }
                ],
                rating: [
                    {
                        required: true,
                        message: '请选择满意度评分',
                        trigger: 'change'
                    }
                ]
            }
        }
    },
    methods: {
        async submitComment() {
            try {
                await this.$refs.commentFormRef.validate();
                this.comments.date = new Date().toLocaleString();
                this.comments.rating = this.comments.rating.toString();
                const response = await this.$axios.post('/api/upload/comment', this.comments).then((response) => {
                    this.submissionMessage = '感谢您的宝贵反馈！';
                    this.isSubmissionSuccess = true;
                }).catch((err) => {
                    alert(err);
                });
                this.$refs.commentFormRef.resetFields();
            } catch (error) {
                console.error('Error sending message:', error);
                this.submissionMessage = error.message || '提交失败，请稍后重试';
                this.isSubmissionSuccess = false;
            }
        }
    }
}
</script>

<style scoped>
.comment-section {
    display: flex;
    justify-content: center;
    align-items: center;
    min-height: 100vh;
    padding: 20px;
}

.comment-container {
    background-color: white;
    border-radius: 16px;
    box-shadow: 0 10px 25px rgba(0, 0, 0, 0.1);
    padding: 40px;
    width: 100%;
    max-width: 600px;
    transition: all 0.3s ease;
}

.section-title {
    text-align: center;
    color: #2c3e50;
    margin-bottom: 30px;
    font-size: 24px;
    font-weight: 600;
    position: relative;
}

.section-title::after {
    content: '';
    position: absolute;
    bottom: -10px;
    left: 50%;
    transform: translateX(-50%);
    width: 80px;
    height: 3px;
    background: linear-gradient(to right, #4CAF50, #2196F3);
}

.comment-form {
    margin-top: 20px;
}

.rating-item {
    display: flex;
    flex-direction: column;
    align-items: center;
    margin: 20px 0;
}

.rating-label {
    color: #34495e;
    margin-bottom: 10px;
    font-size: 16px;
}

.custom-rate {
    font-size: 28px;
}

.submit-button {
    width: 100%;
    padding: 12px;
    border-radius: 10px;
    background: linear-gradient(to right, #4CAF50, #2196F3);
    border: none;
    transition: all 0.3s ease;
}

.submit-button:hover:not(:disabled) {
    transform: translateY(-3px);
    box-shadow: 0 5px 15px rgba(76, 175, 80, 0.4);
}

.submit-button:disabled {
    opacity: 0.6;
    cursor: not-allowed;
}

.submission-alert {
    margin-top: 20px;
}
</style>