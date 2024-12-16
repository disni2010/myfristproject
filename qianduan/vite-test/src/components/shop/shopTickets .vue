<template>
    <div class="main-container">
        <!-- 购票列表 -->
        <div :style="{ display: goupiaoliebiao }" class="ticket-container">
            <h2 class="section-title">购票列表</h2>
            <div class="ticket-list">
                <div v-for="item in this.$store.state.Tickets" :key="item.ticketid" class="ticket-item"
                    @click="showTicketDetails(item)">
                    <h3 class="ticket-name">{{ item.ticketname }}</h3>

                    <div class="info-section">
                        <div class="info-item">
                            <span class="label">价格:</span>
                            <span class="value">¥{{ item.ticketprice }}</span>
                        </div>
                        <div class="info-item">
                            <span class="label">剩余数量:</span>
                            <span class="value">{{ item.ticketnum }}</span>
                        </div>
                        <div class="info-item">
                            <span class="label">开售日期:</span>
                            <span class="value">{{ item.startDate }}</span>
                        </div>
                        <div class="info-item">
                            <span class="label">截止日期:</span>
                            <span class="value">{{ item.endDate }}</span>
                        </div>
                        <template v-for="scen in this.$store.state.Scenicspots" :key="scen.scenicspotid">
                            <div v-if="item.scenicspotid === scen.scenicspotid">
                                <div class="info-item">
                                    <span class="label">景点名称:</span>
                                    <span class="value">{{ scen.scenicspotname }}</span>
                                </div>
                            </div>
                        </template>
                    </div>
                </div>
            </div>
        </div>

        <!-- 购票查询 -->
        <div :style="{ display: goupiaochaxun }" class="order-container">
            <h2 class="section-title">购票查询</h2>
            <div v-if="this.$store.state.mainuser.orders.length > 0" class="order-list">
                <div v-for="item in this.$store.state.mainuser.orders" :key="item.orderid" class="order-item">
                    <div class="info-section">
                        <div class="card-header">
                            <span class="id-tag">订单ID: {{ item.orderid }}</span>
                        </div>
                        <div class="info-item">
                            <span class="label">价格:</span>
                            <span class="value">¥{{ item.orderprice }}</span>
                        </div>
                        <div class="info-item">
                            <span class="label">购票数量:</span>
                            <span class="value">{{ item.ordernum }}</span>
                        </div>
                        <div class="info-item">
                            <span class="label">景区名称:</span>
                            <span class="value">{{ item.orderticketname }}</span>
                        </div>
                        <div class="info-item">
                            <span class="label">票务状态:</span>
                            <span class="value">{{ item.orderstatus === "1" ? "未完成" : item.orderstatus === "2" ? "已完成" :
                                item.orderstatus === "3" ? "退款中" :
                                    "退款完成"
                                }}</span>
                        </div>
                        <div class="info-item">
                            <span class="label">总价:</span>
                            <span class="value">{{ item.ordernum
                                * item.orderprice }}</span>
                        </div>
                        <!-- 支付按钮 为支付的按钮可以按下 -->
                        <div v-if="item.orderstatus === '1'" class="refund-actions">
                            <button @click="PayOrder(item)" class="btn-primary">支付订单</button>
                        </div>
                        <!-- 退款按钮，只有未完成的订单才可以退款 -->
                        <div v-if="item.orderstatus === '2'" class="refund-actions">
                            <button @click="RefundOrder(item)" class="btn-refund">申请退款</button>
                        </div>

                        <!-- 退款取消按钮，只有退款中的订单才可以取消 -->
                        <div v-if="item.orderstatus === '3'" class="refund-actions">
                            <button @click="CancelRefundOrder(item)" class="btn-secondary">取消退款</button>
                        </div>

                        <div v-if="item.orderstatus === '4'" class="refund-actions">
                            <button class="btn-secondary">退款完成</button>
                        </div>
                    </div>
                </div>
            </div>
            <div v-else class="empty-state">
                <div class="empty-icon">🎫</div>
                <h4>没有购买记录</h4>
            </div>
        </div>

        <el-dialog v-model="dialogVisible" title="确认退款" width="500px">
            <el-form :model="refundForm" label-width="80px">
                <el-form-item label="退款原因" required>
                    <el-select v-model="refundForm.reason" placeholder="请选择退款原因" style="width: 100%;">
                        <el-option label="行程变更" value="行程变更"></el-option>
                        <el-option label="景点不满意" value="景点不满意"></el-option>
                        <el-option label="价格太高" value="价格太高"></el-option>
                        <el-option label="其他原因" value="其他原因"></el-option>
                    </el-select>
                </el-form-item>
                <el-form-item label="详细说明" v-if="refundForm.reason === '其他原因'">
                    <el-input v-model="refundForm.description" type="textarea" :rows="3"
                        placeholder="请输入详细退款原因"></el-input>
                </el-form-item>
            </el-form>
            <div>
                <el-button @click="dialogVisible = false">取 消</el-button>
                <el-button type="primary" @click="confirmRefund" :disabled="!refundForm.reason">
                    确认退款
                </el-button>
            </div>
        </el-dialog>

        <!-- 退款查询 -->
        <div :style="{ display: tuikuanchaxun }" class="refund-container">
            <h2 class="section-title">退款查询</h2>
            <div v-if="this.$store.state.Refunds.length > 0" class="refund-list">
                <template v-for="item in this.$store.state.Refunds" :key="item.orderid">
                    <div class="refund-item">
                        <div class="info-section">
                            <div class="card-header">
                                <span class="id-tag">订单ID: {{ item.orderid }}</span>
                            </div>
                            <div class="info-item">
                                <span class="label">价格:</span>
                                <span class="value">¥{{ item.refundprice }}</span>
                            </div>
                            <div class="info-item">
                                <span class="label">数量:</span>
                                <span class="value">{{ item.refundnum }}</span>
                            </div>
                            <div class="info-item">
                                <span class="label">总价:</span>
                                <span class="value">{{ item.refundprice
                                    * item.refundnum }}</span>
                            </div>
                            <div class="info-item">
                                <span class="label">景区名称:</span>
                                <span class="value">{{ item.ticketname }}</span>
                            </div>
                            <div class="info-item">
                                <span class="label">退款状态:</span>
                                <span class="value">{{ item.refundstatus === "1" ? "申请退款中" : item.refundstatus === "2" ?
                                    "退款成功" : item.refundstatus === "3" ? "取消退款" : "退款驳回" }}</span>
                            </div>
                            <div class="info-item">
                                <span class="label">申请日期:</span>
                                <span class="value">{{ item.date }}</span>
                            </div>
                            <div class="info-item">
                                <span class="label">退款原因:</span>
                                <span class="value">{{ item.reason }}</span>
                            </div>
                            <div v-if="item.refundstatus === '4'">
                                <span class="label">驳回原因:</span>
                                <span class="value">{{ item.note }}</span>
                            </div>
                        </div>
                    </div>
                </template>
            </div>
            <div v-else class="empty-state">
                <div class="empty-icon">↩️</div>
                <h4>没有退款处理中</h4>
            </div>
        </div>

        <!-- 票详情弹窗 -->
        <div v-if="selectedTicket" class="modal-overlay">
            <div class="modal-content">
                <h3>票详情</h3>
                <div class="ticket-details">
                    <div class="detail-item">
                        <span class="label">票名称:</span>
                        <span class="value">{{ selectedTicket.ticketname }}</span>
                    </div>
                    <div class="detail-item">
                        <span class="label">票价:</span>
                        <span class="value">¥{{ selectedTicket.ticketprice }}</span>
                    </div>
                    <div class="detail-item">
                        <span class="label">剩余数量:</span>
                        <span class="value">{{ selectedTicket.ticketnum }}</span>
                    </div>
                    <div class="detail-item">
                        <span class="label">开售日期:</span>
                        <span class="value">{{ selectedTicket.startDate }}</span>
                    </div>
                    <div class="detail-item">
                        <span class="label">截止日期:</span>
                        <span class="value">{{ selectedTicket.endDate }}</span>
                    </div>
                    <!-- 景点信息 -->
                    <template v-for="scen in this.$store.state.Scenicspots" :key="scen.scenicspotid">
                        <div v-if="selectedTicket.scenicspotid === scen.scenicspotid" class="detail-item">
                            <span class="label">景点名称:</span>
                            <span class="value">{{ scen.scenicspotname }}</span>
                        </div>
                    </template>
                    <!-- 添加数量选择器 -->
                    <div class="detail-item quantity-selector">
                        <span class="label">购买数量:</span>
                        <div class="quantity-controls">
                            <button class="quantity-btn" @click="decreaseQuantity"
                                :disabled="selectedQuantity <= 1">-</button>
                            <input type="number" v-model.number="selectedQuantity" :min="1"
                                :max="selectedTicket.ticketnum" class="quantity-input" @input="validateQuantity">
                            <button class="quantity-btn" @click="increaseQuantity"
                                :disabled="selectedQuantity >= selectedTicket.ticketnum">+</button>
                        </div>
                    </div>
                    <!-- 显示总价 -->
                    <div class="detail-item total-price">
                        <span class="label">总价:</span>
                        <span class="value price-highlight">¥{{ totalPrice }}</span>
                    </div>
                </div>
                <div class="button-group">
                    <button class="btn-primary" @click="buyTicket"
                        :disabled="selectedQuantity > selectedTicket.ticketnum">购买</button>
                    <button class="btn-secondary" @click="closeDetails">关闭</button>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
import axios from 'axios';

export default {
    data() {
        return {
            goupiaoliebiao: "none",
            goupiaochaxun: "none",
            tuikuanchaxun: "none",
            selectedTicket: null,
            isRefund: false,
            selectedQuantity: 1, // 添加数量选择变量
            msg: {
                field1: "",
                field2: "",
                field3: ""
            },
            dialogVisible: false,
            currentOrder: null,
            refundForm: {
                reason: '',
                description: ''
            },
            userid: ""
        };
    },
    computed: {
        totalPrice() {
            if (!this.selectedTicket) return 0;
            return (this.selectedTicket.ticketprice * this.selectedQuantity).toFixed(2);
        }
    },
    mounted() {
        this.init()
        console.log("userid:", this.$store.state.mainuser.user.userid);
        console.log("storeuserid:", this.$store.state.mainuser.user.userid);
    }
    ,
    watch: {
        '$route': 'init'
    },
    methods: {
        init() {
            console.log("订单:", this.$store.state.mainuser.orders);
            this.goupiaoliebiao = "none";
            this.goupiaochaxun = "none";
            this.tuikuanchaxun = "none";
            const method = this.$route.params.method;
            switch (method) {
                case "goupiaoliebiao":
                    this.goupiaoliebiao = "block";
                    break;
                case "goupiaochaxun":
                    this.goupiaochaxun = "block";
                    break;
                case "tuikuanchaxun":
                    this.tuikuanchaxun = "block";
                    break;
            }
        },
        showTicketDetails(ticket) {
            this.selectedTicket = ticket;
            this.selectedQuantity = 1; // 重置数量选择
        },
        buyTicket() {
            if (this.selectedQuantity <= 0 || this.selectedQuantity > this.selectedTicket.ticketnum) {
                alert('请选择有效的购票数量');
                return;
            }
            var ticket = {
                ticketname: this.selectedTicket.ticketname,
                ticketprice: this.selectedTicket.ticketprice,
                ticketnum: this.selectedQuantity.toString(),
                scenicspotid: this.selectedTicket.scenicspotid,
            }
            this.msg.field1 = JSON.stringify(this.$store.state.mainuser.user.userid)
            this.msg.field2 = JSON.stringify(ticket)

            axios.post('/api/upload/buytickets', this.msg)
                .then(response => {
                    // 成功响应
                    alert("购买成功");
                    console.log(response);
                    this.$store.commit('setOrdersList', response.data.msg);
                    const index = this.$store.state.Tickets.findIndex(t => t.ticketname === ticket.ticketname);
                    console.log("下标：" + index);
                    if (index !== -1) {
                        this.$store.state.Tickets[index].ticketnum = parseInt(this.$store.state.Tickets[index].ticketnum) - parseInt(ticket.ticketnum);
                        this.$store.commit('setTicketsList', this.$store.state.Tickets);
                        this.$store.dispatch('saveTicketsListToStorage');
                        location.reload();
                    }
                    if (response.data.msg.code != 500) {
                        //向后端传递购票数据进行购票
                        alert(`您已购买 ${this.selectedTicket.ticketname
                            } ${this.selectedQuantity} 张，总价：¥${this.totalPrice} `);
                    }

                })
                .catch(error => {
                    // 错误响应
                    if (error.response && error.response.data) {
                        alert("购买失败:" + error.response.data.msg);
                    }
                    console.log("失败:" + error);
                });
            this.closeDetails();
        },
        PayOrder(item) {
            this.msg.field1 = JSON.stringify(this.$store.state.mainuser.user.userid)
            this.msg.field2 = JSON.stringify(item)
            console.log("发送的包");
            console.log(this.msg);
            axios.post('/api/upload/payorder', this.msg)
                .then(response => {
                    // 成功响应
                    alert("支付成功");
                    console.log(response);
                    this.$store.commit('setOrdersList', response.data.msg);
                    this.$store.dispatch('saveOrdersListToStorage');
                    location.reload();
                })
                .catch(error => {
                    // 错误响应
                    if (error.response && error.response.data) {
                        alert("支付失败:" + error.response.data.msg);
                    }
                    console.log("其他失败:" + error);
                })
        },
        CancelRefundOrder(item) {
            const comfin = confirm("是否确认取消退款？");
            if (comfin) {

                var refund = {
                    orderid: item.orderid,
                    refunduserid: item.orderuserid
                }

                this.msg.field1 = JSON.stringify(this.$store.state.mainuser.user.userid)
                this.msg.field2 = JSON.stringify(refund)
                console.log("发送的包");
                console.log(this.msg);
                axios.post('/api/upload/cancelrefundorder', this.msg)
                    .then(response => {
                        // 成功响应
                        alert("取消退款成功");
                        this.$store.commit('setRefundsList', response.data.msg);
                        this.$store.dispatch('saveRefundsListToStorage');
                        response.data.msg.refunds.forEach((refund) => {
                            if (refund.refundstatus === "3") {
                                this.$store.state.mainuser.orders.forEach((order, orderindex) => {
                                    if (refund.orderid === order.orderid) {
                                        this.$store.state.mainuser.orders[orderindex].orderstatus = "2"
                                    }
                                });
                            }
                        });
                        this.$store.dispatch('saveOrdersListToStorage');
                        console.log(response.data);
                        location.reload();
                    })
                    .catch(error => {
                        // 错误响应
                        if (error.response && error.response.data) {
                            alert("取消失败:" + error.response.data.msg);
                        }
                        console.log("其他失败:" + error);
                    })
            }
        },
        closeDetails() {
            this.selectedTicket = null;
            this.selectedQuantity = 1; // 重置数量选择
        },
        // 数量控制方法
        decreaseQuantity() {
            if (this.selectedQuantity > 1) {
                this.selectedQuantity--;
            }
        },
        increaseQuantity() {
            if (this.selectedQuantity < this.selectedTicket.ticketnum) {
                this.selectedQuantity++;
            }
        },
        validateQuantity() {
            // 确保输入的数量在有效范围内
            if (this.selectedQuantity < 1) {
                this.selectedQuantity = 1;
            } else if (this.selectedQuantity > this.selectedTicket.ticketnum) {
                this.selectedQuantity = this.selectedTicket.ticketnum;
            }
            // 确保输入为整数
            this.selectedQuantity = Math.floor(this.selectedQuantity);
        },
        RefundOrder(order) {
            this.currentOrder = order;
            this.dialogVisible = true;
            // 重置表单
            this.refundForm = {
                reason: '',
                description: ''
            };
        },
        confirmRefund() {
            if (!this.refundForm.reason) {
                this.$message.error('请选择退款原因');
                return;
            }

            // 准备退款数据
            // const obj1 = { x: 1, y: 2 };
            // const obj2 = { ...obj1, z: 3 };
            // console.log(obj2); // { x: 1, y: 2, z: 3 }
            const refundData = {
                ...this.currentOrder,
                refundReason: this.refundForm.reason === '其他原因'
                    ? this.refundForm.description
                    : this.refundForm.reason
            };
            //退款数据 Refunds orderid(订单id) refundprice(退款价格) refundnum(退款数量) refunduserid(退款用户id) refundstatus(退款状态:1申请退款2退款成功3退款失败) ticketname(票务名称) Data(日期) reason(退款原因) Note(备注)
            var Refunds = {
                refunduserid: this.$store.state.mainuser.user.userid,
                reason: refundData.refundReason,
            }

            this.msg.field1 = JSON.stringify(this.$store.state.mainuser.user.userid)
            this.msg.field2 = JSON.stringify(refundData)
            this.msg.field3 = JSON.stringify(Refunds)
            console.log("发送的包");
            console.log(this.msg);
            // 调用退款接口
            this.$axios.post('/api/upload/refundorder', this.msg)
                .then(response => {
                    console.log(response);
                    this.$message.success('退款申请提交成功');
                    this.$store.commit('setRefundsList', response.data.msg);
                    this.$store.dispatch('saveRefundsListToStorage');
                    console.log(Array.isArray(response.data.msg.refunds));
                    response.data.msg.refunds.forEach((refund) => {
                        if (refund.refundstatus === 1) {
                            this.$store.state.mainuser.orders.forEach((order, orderindex) => {
                                if (refund.orderid === order.orderid) {
                                    this.$store.state.mainuser.orders[orderindex].orderstatus = "3"
                                }
                            });
                        }
                    });
                    this.$store.dispatch('saveOrdersListToStorage');
                    this.dialogVisible = false;
                    location.reload();
                })
                .catch(error => {
                    if (error.response && error.response.data) {
                        this.$message.error(`退款失败: ${error.response.data.msg}`);
                    }
                    console.log("退款失败:", error);
                });
        }
    },
};
</script>

<style scoped>
.main-container {
    background-color: #f5f5f5;
    padding: 0px;
    /* 移除容器内边距 */
    font-family: system-ui, -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
    width: 100%;
    /* 确保容器占满宽度 */
    height: 100%;
    margin: 0px;
    position: relative;
}

.section-title {
    font-size: 24px;
    color: #333;
    margin-bottom: 20px;
    padding-bottom: 10px;
    border-bottom: 3px solid #3b82f6;
    display: inline-block;
}

.ticket-list,
.order-list,
.refund-list {
    display: flex;
    flex-wrap: wrap;
    gap: 20px;
    margin-bottom: 40px;
    width: 100%;
    /* 确保列表占满容器宽度 */
}

.ticket-container,
.order-container,
.refund-container {
    padding: 20px 40px;
    /* 增加左右内边距 */
    width: 100%;
    background-color: #f5f5f5;
}

.ticket-item,
.order-item,
.refund-item {
    flex: 1;
    min-width: calc(22% - 15px);
    /* 略微增加最小宽度 */
    max-width: calc(22% - 15px);
    background: white;
    border-radius: 10px;
    padding: 20px;
    box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
    transition: transform 0.3s ease, box-shadow 0.3s ease;
    cursor: pointer;
    margin: 0;
    /* 移除外边距 */
}

.ticket-item:hover {
    transform: translateY(-5px);
    box-shadow: 0 4px 8px rgba(0, 0, 0, 0.15);
}

.card-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
}

.id-tag {
    background: #e5e7eb;
    padding: 4px 8px;
    border-radius: 4px;
    font-size: 14px;
    color: #374151;
}

.price-tag {
    font-size: 20px;
    font-weight: bold;
    color: #3b82f6;
}

.status-tag {
    background: #fee2e2;
    color: #dc2626;
    padding: 4px 12px;
    border-radius: 15px;
    font-size: 14px;
    font-weight: 500;
}

.ticket-name {
    font-size: 18px;
    color: #1f2937;
    margin-bottom: 15px;
    font-weight: 500;
}

.info-section {
    display: flex;
    flex-direction: column;
    gap: 10px;
}

.info-item {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 8px;
    background: #f8fafc;
    border-radius: 6px;
}

.label {
    color: #6b7280;
    font-size: 14px;
}

.value {
    font-weight: 500;
    color: #374151;
}

.empty-state {
    width: 100%;
    text-align: center;
    padding: 40px;
    background: white;
    border-radius: 10px;
    box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.empty-icon {
    font-size: 40px;
    margin-bottom: 10px;
}

.modal-overlay {
    position: fixed;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    background: rgba(0, 0, 0, 0.5);
    display: flex;
    align-items: center;
    justify-content: center;
    z-index: 1000;
}

.modal-content {
    background: white;
    padding: 30px;
    border-radius: 15px;
    width: 90%;
    max-width: 400px;
    text-align: center;
}

.modal-content h3 {
    font-size: 22px;
    color: #1f2937;
    margin-bottom: 15px;
}

.modal-content p {
    color: #4b5563;
    margin-bottom: 20px;
}

.button-group {
    display: flex;
    gap: 10px;
    justify-content: center;
}

.btn-primary,
.btn-secondary,
.btn-refund {
    padding: 10px 20px;
    border-radius: 6px;
    font-weight: 500;
    cursor: pointer;
    transition: background-color 0.3s ease;
}

.btn-primary {
    background: #3b82f6;
    color: white;
    border: none;
}

.btn-primary:hover {
    background: #2563eb;
}

.btn-secondary {
    background: #e5e7eb;
    color: #4b5563;
    border: none;
}

.btn-secondary:hover {
    background: #d1d5db;
}

.btn-refund {
    background: #ffb310;
    color: #4b5563;
    border: none;
}

.btn-refund:hover {
    background: #cf0f0f;
}

.modal-overlay {
    position: fixed;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    background-color: rgba(0, 0, 0, 0.5);
    display: flex;
    justify-content: center;
    align-items: center;
    z-index: 1000;
}

.modal-content {
    background-color: white;
    padding: 2rem;
    border-radius: 8px;
    min-width: 300px;
    max-width: 500px;
}

.ticket-details {
    margin: 1.5rem 0;
}

.detail-item {
    margin: 0.5rem 0;
    display: flex;
    justify-content: space-between;
    align-items: center;
}

.detail-item .label {
    font-weight: bold;
    color: #666;
}

.detail-item .value {
    color: #333;
}

.button-group {
    display: flex;
    justify-content: flex-end;
    gap: 1rem;
    margin-top: 1.5rem;
}

.btn-primary,
.btn-secondary {
    padding: 0.5rem 1rem;
    border-radius: 4px;
    cursor: pointer;
    border: none;
}

.btn-primary {
    background-color: #4CAF50;
    color: white;
}

.btn-secondary {
    background-color: #f5f5f5;
    color: #333;
}

.quantity-selector {
    margin: 15px 0;
}

.quantity-controls {
    display: flex;
    align-items: center;
    gap: 8px;
}

.quantity-btn {
    width: 30px;
    height: 30px;
    border: 1px solid #ddd;
    background: #f5f5f5;
    border-radius: 4px;
    cursor: pointer;
    display: flex;
    align-items: center;
    justify-content: center;
    font-size: 18px;
    transition: all 0.3s ease;
}

.quantity-btn:disabled {
    opacity: 0.5;
    cursor: not-allowed;
}

.quantity-btn:hover:not(:disabled) {
    background: #e0e0e0;
}

.quantity-input {
    width: 60px;
    height: 30px;
    text-align: center;
    border: 1px solid #ddd;
    border-radius: 4px;
    font-size: 14px;
}

.quantity-input::-webkit-inner-spin-button,
.quantity-input::-webkit-outer-spin-button {
    -webkit-appearance: none;
    margin: 0;
}

.total-price {
    margin-top: 15px;
    border-top: 1px solid #eee;
    padding-top: 15px;
}

.price-highlight {
    color: #f43f5e;
    font-size: 1.25rem;
    font-weight: bold;
}


/* 响应式布局调整
@media (max-width: 1400px) {
    .ticket-item,
    .order-item,
    .refund-item {
        max-width: calc(33.33% - 14px);
    }
}

@media (max-width: 1024px) {
    .ticket-container,
    .order-container,
    .refund-container {
        padding: 20px;
    }
    
    .ticket-item,
    .order-item,
    .refund-item {
        max-width: calc(50% - 10px);
    }
}

@media (max-width: 640px) {
    .ticket-item,
    .order-item,
    .refund-item {
        max-width: 100%;
    }
} */
</style>