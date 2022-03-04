                                                <template>
                                                <v-container>
                                                <v-row>
                                                <v-col>
                                                <v-container class="container1">

                                                <v-col cols="12" sm="6" class="my-n4">
                                                <v-select :items="items_character" label="サポートキャラを選択" outlined v-model="selected_character"></v-select>
                                                </v-col>
                                                <v-col cols="12" sm="6" class="my-n4">
                                                <v-select :items="items_support" label="難易度を選択してください" outlined v-model="selected_support"></v-select>
                                                </v-col>
                                                <v-col cols="12" sm="6" class="my-n4">
                                                <v-select :items="num_of_tornament" label="何回戦かを選択してください" outlined v-model="selected_num_of_tornament"></v-select>
                                                </v-col> 

                                                <v-col cols="12" sm="6" class="my-n4">
                                                <v-select :items="items_gokigen" label="サポートキャラのご機嫌度" outlined v-model="selected_gokigen"></v-select>
                                                </v-col>  
                                                <v-col cols="12" sm="6" class="my-n4">
                                                <v-select :items="lost_point" label="失点数を入力してください" outlined v-model="selected_lost_point"></v-select>
                                                </v-col>
                                                <v-col cols="12" sm="6" class="my-n4">
                                                <v-select :items="goal_difference" label="得失点差を入力してください" outlined v-model="selected_goal_difference"></v-select>
                                                </v-col>                   

                                                <v-tabs v-model="tab" grow class="mb-4">
                                                <v-tab>野手経験値</v-tab>
                                                <v-tab>投手経験値</v-tab>
                                                </v-tabs>
                                                <v-tabs-items v-model="tab">
                                                <v-tab-item>
                                                <v-row>
                                                <v-col cols="5"><label>ヒット</label></v-col>
                                                <v-col cols="2"><v-text-field v-model="count_hit"></v-text-field></v-col>
                                                <v-col cols="2"><v-btn color="accent" elevation="2" outlined v-on:click="counter('hit','+')"><v-icon dark>mdi-plus</v-icon></v-btn></v-col>
                                                <v-col cols="2"><v-btn color="accent" elevation="2" outlined v-on:click="counter('hit','-')"><v-icon dark>mdi-minus</v-icon></v-btn></v-col>
                                                </v-row>
                                                <v-row>
                                                <v-col cols="5"><label>ツーベース</label></v-col>
                                                <v-col cols="2"><v-text-field v-model="count_twobase"></v-text-field></v-col>
                                                <v-col cols="2"><v-btn color="accent" elevation="2" outlined v-on:click="counter('twobase','+')"><v-icon dark>mdi-plus</v-icon></v-btn></v-col>
                                                <v-col cols="2"><v-btn color="accent" elevation="2" outlined v-on:click="counter('twobase','-')"><v-icon dark>mdi-minus</v-icon></v-btn></v-col>
                                                </v-row>
                                                <v-row>
                                                <v-col cols="5"><label>スリーベース</label></v-col>
                                                <v-col cols="2"><v-text-field v-model="count_threebase"></v-text-field></v-col>
                                                <v-col cols="2"><v-btn color="accent" elevation="2" outlined v-on:click="counter('threebase','+')"><v-icon dark>mdi-plus</v-icon></v-btn></v-col>
                                                <v-col cols="2"><v-btn color="accent" elevation="2" outlined v-on:click="counter('threebase','-')"><v-icon dark>mdi-minus</v-icon></v-btn></v-col>
                                                </v-row>
                                                <v-row>
                                                <v-col cols="5"><label>ホームラン</label></v-col>
                                                <v-col cols="2"><v-text-field v-model="count_homerun"></v-text-field></v-col>
                                                <v-col cols="2"><v-btn color="accent" elevation="2" outlined v-on:click="counter('homerun','+')"><v-icon dark>mdi-plus</v-icon></v-btn></v-col>
                                                <v-col cols="2"><v-btn color="accent" elevation="2" outlined v-on:click="counter('homerun','-')"><v-icon dark>mdi-minus</v-icon></v-btn></v-col>
                                                </v-row>
                                                <v-row>
                                                <v-col cols="5"><label>バント</label></v-col>
                                                <v-col cols="2"><v-text-field v-model="count_bunt"></v-text-field></v-col>
                                                <v-col cols="2"><v-btn color="accent" elevation="2" outlined v-on:click="counter('bunt','+')"><v-icon dark>mdi-plus</v-icon></v-btn></v-col>
                                                <v-col cols="2"><v-btn color="accent" elevation="2" outlined v-on:click="counter('bunt','-')"><v-icon dark>mdi-minus</v-icon></v-btn></v-col>
                                                </v-row>
                                                <v-row>
                                                <v-col cols="5"><label>犠牲フライ</label></v-col>
                                                <v-col cols="2"><v-text-field v-model="count_fly"></v-text-field></v-col>
                                                <v-col cols="2"><v-btn color="accent" elevation="2" outlined v-on:click="counter('fly','+')"><v-icon dark>mdi-plus</v-icon></v-btn></v-col>
                                                <v-col cols="2"><v-btn color="accent" elevation="2" outlined v-on:click="counter('fly','-')"><v-icon dark>mdi-minus</v-icon></v-btn></v-col>
                                                </v-row>
                                                <v-row>
                                                <v-col cols="5"><label>盗塁</label></v-col>
                                                <v-col cols="2"><v-text-field v-model="count_steal"></v-text-field></v-col>
                                                <v-col cols="2"><v-btn color="accent" elevation="2" outlined v-on:click="counter('steal','+')"><v-icon dark>mdi-plus</v-icon></v-btn></v-col>
                                                <v-col cols="2"><v-btn color="accent" elevation="2" outlined v-on:click="counter('steal','-')"><v-icon dark>mdi-minus</v-icon></v-btn></v-col>
                                                </v-row>
                                                </v-tab-item>

                                                <v-tab-item>
                                                <v-row>
                                                <v-col cols="5"><label>登板イニング毎</label></v-col>
                                                <v-col cols="2"><v-text-field v-model="count_inings"></v-text-field></v-col>
                                                <v-col cols="2"><v-btn color="accent" elevation="2" outlined v-on:click="counter('inings','+')"><v-icon dark>mdi-plus</v-icon></v-btn></v-col>
                                                <v-col cols="2"><v-btn color="accent" elevation="2" outlined v-on:click="counter('inings','-')"><v-icon dark>mdi-minus</v-icon></v-btn></v-col>
                                                </v-row>
                                                <v-row>
                                                <v-col cols="5"><label>ストレート系で奪三振</label></v-col>
                                                <v-col cols="2"><v-text-field v-model="count_sunsinwithfastball"></v-text-field></v-col>
                                                <v-col cols="2"><v-btn color="accent" elevation="2" outlined v-on:click="counter('sunsinwithfastball','+')"><v-icon dark>mdi-plus</v-icon></v-btn></v-col>
                                                <v-col cols="2"><v-btn color="accent" elevation="2" outlined v-on:click="counter('sunsinwithfastball','-')"><v-icon dark>mdi-minus</v-icon></v-btn></v-col>
                                                </v-row>
                                                <v-row>
                                                <v-col cols="5"><label>ストレート系でゴロアウト</label></v-col>
                                                <v-col cols="2"><v-text-field v-model="count_gorowithfastball"></v-text-field></v-col>
                                                <v-col cols="2"><v-btn color="accent" elevation="2" outlined v-on:click="counter('gorowithfastball','+')"><v-icon dark>mdi-plus</v-icon></v-btn></v-col>
                                                <v-col cols="2"><v-btn color="accent" elevation="2" outlined v-on:click="counter('gorowithfastball','-')"><v-icon dark>mdi-minus</v-icon></v-btn></v-col>
                                                </v-row>
                                                <v-row>
                                                <v-col cols="5"><label>ストレート系でフライアウト</label></v-col>
                                                <v-col cols="2"><v-text-field v-model="count_flywithfastball"></v-text-field></v-col>
                                                <v-col cols="2"><v-btn color="accent" elevation="2" outlined v-on:click="counter('flywithfastball','+')"><v-icon dark>mdi-plus</v-icon></v-btn></v-col>
                                                <v-col cols="2"><v-btn color="accent" elevation="2" outlined v-on:click="counter('flywithfastball','-')"><v-icon dark>mdi-minus</v-icon></v-btn></v-col>
                                                </v-row>
                                                <v-row>
                                                <v-col cols="5"><label>変化球で奪三振</label></v-col>
                                                <v-col cols="2"><v-text-field v-model="count_sunsinwithbreakingball"></v-text-field></v-col>
                                                <v-col cols="2"><v-btn color="accent" elevation="2" outlined v-on:click="counter('sunsinwithbreakingball','+')"><v-icon dark>mdi-plus</v-icon></v-btn></v-col>
                                                <v-col cols="2"><v-btn color="accent" elevation="2" outlined v-on:click="counter('sunsinwithbreakingball','-')"><v-icon dark>mdi-minus</v-icon></v-btn></v-col>
                                                </v-row>
                                                <v-row>
                                                <v-col cols="5"><label>変化球でゴロアウト</label></v-col>
                                                <v-col cols="2"><v-text-field v-model="count_gorowithbreakingball"></v-text-field></v-col>
                                                <v-col cols="2"><v-btn color="accent" elevation="2" outlined v-on:click="counter('gorowithbreakingball','+')"><v-icon dark>mdi-plus</v-icon></v-btn></v-col>
                                                <v-col cols="2"><v-btn color="accent" elevation="2" outlined v-on:click="counter('gorowithbreakingball','+')"><v-icon dark>mdi-minus</v-icon></v-btn></v-col>
                                                </v-row>
                                                <v-row>
                                                <v-col cols="5"><label>変化球でフライアウト</label></v-col>
                                                <v-col cols="2"><v-text-field v-model="count_flywithbreakingball"></v-text-field></v-col>
                                                <v-col cols="2"><v-btn color="accent" elevation="2" outlined v-on:click="counter('flywithbreakingball','+')"><v-icon dark>mdi-plus</v-icon></v-btn></v-col>
                                                <v-col cols="2"><v-btn color="accent" elevation="2" outlined v-on:click="counter('flywithbreakingball','-')"><v-icon dark>mdi-minus</v-icon></v-btn></v-col>
                                                </v-row>
                                                <v-row>
                                                <v-col cols="5"><label>併殺</label></v-col>
                                                <v-col cols="2"><v-text-field v-model="count_gettsu"></v-text-field></v-col>
                                                <v-col cols="2"><v-btn color="accent" elevation="2" outlined v-on:click="counter('gettsu','+')"><v-icon dark>mdi-plus</v-icon></v-btn></v-col>
                                                <v-col cols="2"><v-btn color="accent" elevation="2" outlined v-on:click="counter('gettsu','-')"><v-icon dark>mdi-minus</v-icon></v-btn></v-col>
                                                </v-row>
                                                </v-tab-item>
                                                </v-tabs-items>
                                                </v-container>
                                                </v-col>
                                                <v-col>                                
                                                <v-row justify="center">
                                                <v-btn color="red accent-2" elevation="2" outlined v-on:click="reset">リセット</v-btn>
                                </v-row>
                                                <v-row>
                                                <div class="balloon6">
                                                <div class="faceicon">
                                                <img src="../assets/yabe.png" width="150px" height="150px">
                                                </div>
                                                <div class="chatting">
                                                <div class="says">
                                                <p>{{ yabes_message }}</p>
                                                </div>
                                                </div>
                                                </div>
                                                </v-row>
                                                <v-row>
                                                <v-container>
                                                <table border="1">
                                      <tr><th v-for="thead in theads" v-bind:key="thead">{{thead}}</th></tr>
                                                <tr><td v-for="tdata in goukei" v-bind:key="tdata">{{tdata}}</td></tr>
                                                </table>
                                                </v-container>
                                                </template>