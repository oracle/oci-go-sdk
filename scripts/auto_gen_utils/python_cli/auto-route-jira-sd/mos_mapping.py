import pandas as pd
import constants

PATH_TO_MOS_MAP_CSV = constants.PATH_TO_CSV + "mos_mapping_file.csv"


def load_mos_mapping_csv():
    df = pd.read_csv(constants.PATH_TO_MOS_MAP_CSV, encoding='latin-1')
    return df


def fetch_comp_sub_comp(jira_q):
    """
    :param jira_q: jira sd queue code
    :return: list of mos component code and subcomponent code
    """
    df = load_mos_mapping_csv()
    df = df.query('`JIRA_PROJECT` == "' + jira_q + '"')
    if not df.empty:
        component_code = df["COMPONENT_CODE"].values.tolist()[0]
        sub_comp_code = df["SUBCOMPONENT_CODE"].values.tolist()[0]
    else:
        print("Entry for " + jira_q + " not found")
    return [component_code, sub_comp_code]
