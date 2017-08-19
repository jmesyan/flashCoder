import React from 'react';
import { Field, Form, Errors, actions } from 'react-redux-form';
const  regexps = {
	second:/^(((\d|([1-5]\d))|(\d|([1-5]\d))[-,](\d|([1-5]\d)))|(\*\/([1-9]|([1-5]\d)))|(\*))$/,
	minute:/^(((\d|([1-5]\d))|(\d|([1-5]\d))[-,](\d|([1-5]\d)))|(\*\/([1-9]|([1-5]\d)))|(\*))$/,
	hour:/^((([0-9]|(1[0-9]|2[0-3]))|([0-9]|(1[0-9]|2[0-3]))[-,]([0-9]|(1[0-9]|2[0-3])))|(\*\/([1-9]|(1[0-9]|2[0-3])))|(\*))$/,
	day:/^((([1-9]|([12]\d|3[01]))|([1-9]|([12]\d|3[01]))[-,]([1-9]|([12]\d|3[01])))|(\*\/([1-9]|([12]\d|3[01])))|([\*?]))$/,
	month:/^((([1-9]|(1[0-2]))|([1-9]|(1[0-2]))[-,]([1-9]|(1[0-2]))?)|(\*\/([1-9]|(1[0-2]?)))|(\*))$/,
	week:/^(([0-6]|[0-6][-,][0-6])|(\*\/[1-6])|([\*?]))$/,
	dayExist:/^((([1-9]|([12]\d|3[01]))|([1-9]|([12]\d|3[01]))[-,]([1-9]|([12]\d|3[01])))|(\*\/([1-9]|([12]\d|3[01]))))$/,
	weekExist:/^(([0-6]|[0-6][-,][0-6])|(\*\/[1-6]))$/,
	period:/^(\d+)-(\d+)$/
}

const isRequired = (val) => val && val.length > 0;
const minMaxValid = (val) => {
	var arr = val.split("-");
	if (parseInt(arr[0]) > parseInt(arr[1])) return false;
	return true;
}



var CronFrom = React.createClass({

	render: function() {
        const {handleSubmit, cronItem} = this.props;
		return (
			<div>
                <Form 
						model="cronForm" 
							validators={{
								'': {
									dayOrWeek:(vals) => {
										if (vals.Day != "" && vals.Week != "" && regexps.dayExist.test(vals.Day) && regexps.weekExist.test(vals.Week)) return false;
										return true;
									}
								},
							}}
							onSubmit = {handleSubmit}
						>
							<div className="form-inline clear" >
								<div className="form-group fl">
									<label className="form-control">任务ID:{cronItem.task.Tid}</label>
								</div>
								<div className="form-group fl ml-20">
									<label className="form-control">任务名称:{cronItem.task.Tname}</label>
								</div>
							</div>
							<div className="form-inline clear mt-20">
								<div className="form-group">
									<Field model="cronForm.Second" validators={{ isRequired, minMaxValid, regValid:(val)=>val == '' || regexps.second.test(val)}}>
											<input type="text" className="form-control fc-p64" placeholder="请输入秒"/>
											<Errors
												wrapper="small"
												className="help-block  form-error"
												show={{ touched: true, focus: false }}
												model="cronForm.Second"
												messages={{
													isRequired: '时间秒不能为空',
													minMaxValid:'最小值不能大于最大值',
													regValid:"时间秒无效"
												}}
											
											/>
									</Field>
							  	</div>
									<div className="form-group">
										<Field model="cronForm.Minute"  validators={{ isRequired, minMaxValid, regValid:(val)=>val == '' || regexps.minute.test(val)}}>
												<input type="text" className="form-control fc-p64" placeholder="请输入分"/>
												<Errors
													wrapper="small"
													className="help-block  form-error"
													show={{ touched: true, focus: false }}
													model="cronForm.Minute"
													messages={{
														isRequired: '时间分不能为空',
														minMaxValid:'最小值不能大于最大值',
														regValid:"时间分无效"
													}}
												
												/>
										</Field>
								</div>
									<div className="form-group">
										<Field model="cronForm.Hour" validators={{ isRequired, minMaxValid, regValid:(val)=>val == '' || regexps.hour.test(val)}}>
												<input type="text" className="form-control fc-p64" placeholder="请输入小时"/>
												<Errors
													wrapper="small"
													className="help-block  form-error"
													show={{ touched: true, focus: false }}
													model="cronForm.Hour"
													messages={{
														isRequired: '时间小时不能为空',
														minMaxValid:'最小值不能大于最大值',
														regValid:"时间小时无效"
													}}
												
												/>
										</Field>
								</div>
									<div className="form-group">
										<Field model="cronForm.Day"  validators={{ isRequired, minMaxValid, regValid:(val)=>val == '' || regexps.day.test(val)}}>
												<input type="text" className="form-control fc-p64" placeholder="请输入月天"/>
												<Errors
													wrapper="small"
													className="help-block  form-error"
													show={{ touched: true, focus: false }}
													model="cronForm.Day"
													messages={{
														isRequired: '月天不能为空',
														minMaxValid:'最小值不能大于最大值',
														regValid:"月天无效",
													}}
												
												/>
												<Errors
													wrapper="small"
													className="help-block  form-error"
													show={{ touched: true, focus: false }}
													model="cronForm"
													messages={{
														dayOrWeek:"月天和周天只能选择一个"
													}}
												
												/>
										</Field>
								</div>
									<div className="form-group">
										<Field model="cronForm.Month"  validators={{ isRequired, minMaxValid, regValid:(val)=>val == '' || regexps.month.test(val)}}>
												<input type="text" className="form-control fc-p64" placeholder="请输入月份"/>
												<Errors
													wrapper="small"
													className="help-block  form-error"
													show={{ touched: true, focus: false }}
													model="cronForm.Month"
													messages={{
														isRequired: '月份不能为空',
														minMaxValid:'最小值不能大于最大值',
														regValid:"月份无效"
													}}
												
												/>
										</Field>
								</div>
									<div className="form-group">
										<Field model="cronForm.Week"   validators={{ isRequired, minMaxValid, regValid:(val)=>val == '' || regexps.week.test(val)}}>
												<input type="text" className="form-control fc-p64" placeholder="请输入周天"/>
												<Errors
													wrapper="small"
													className="help-block  form-error"
													show={{ touched: true, focus: false }}
													model="cronForm.Week"
													messages={{
														isRequired: '周天不能为空',
														minMaxValid:'最小值不能大于最大值',
														regValid:"周天无效",
													}}
												
												/>
												<Errors
													wrapper="small"
													className="help-block  form-error"
													show={{ touched: true, focus: false }}
													model="cronForm"
													messages={{
														dayOrWeek:"月天和周天只能选择一个"
													}}
												
												/>
										</Field>
								</div>
							</div>
							<div className="form-group mt-20">
								<input type="submit" className="btn btn-primary" value="提交任务" />
							</div>	
					</Form>
            </div>
		);
	}

});

module.exports = CronFrom;