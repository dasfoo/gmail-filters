package filter

/*<apps:property name='from' value='asasa'/>
<apps:property name='to' value='asasa'/>
<apps:property name='subject' value='aaa'/>
<apps:property name='hasTheWord' value='aaa'/>
<apps:property name='doesNotHaveTheWord' value='bbb'/>
<apps:property name='hasAttachment' value='true'/>
<apps:property name='label' value='Путешествие'/>
<apps:property name='shouldMarkAsRead' value='true'/>
<apps:property name='shouldStar' value='true'/>
<apps:property name='shouldTrash' value='true'/>
<apps:property name='shouldNeverSpam' value='true'/>
<apps:property name='shouldAlwaysMarkAsImportant' value='true'/>
<apps:property name='sizeOperator' value='s_sl'/>
<apps:property name='sizeUnit' value='s_smb'/>
<apps:property name='subject' value='dddd'/>
<apps:property name='hasAttachment' value='true'/>
<apps:property name='excludeChats' value='true'/>
<apps:property name='shouldNeverMarkAsImportant' value='true'/>
<apps:property name='smartLabelToApply' value='^smartlabel_personal'/>
<apps:property name='size' value='3'/>
<apps:property name='sizeOperator' value='s_sl'/>
<apps:property name='sizeUnit' value='s_smb'/>
*/

type Property struct {
	Name  string `xml:"name,attr"`
	Value string `xml:"value,attr"`
}
