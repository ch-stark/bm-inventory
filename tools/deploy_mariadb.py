import os
import utils
import argparse
import yaml

parser = argparse.ArgumentParser()
parser.add_argument('--service', help='Service name to use', type=str, default=None)
parser.add_argument('--namespace', help='Namespace to use', type=str, default=None)
args = parser.parse_args()

NAMESPACE = args.namespace or 'assisted-installer'


def main():
    src_file = os.path.join(os.getcwd(), "deploy/mariadb/mariadb-configmap.yaml")
    dst_file = os.path.join(os.getcwd(), "build/mariadb-configmap.yaml")
    with open(src_file, "r") as src:
        with open(dst_file, "w+") as dst:
            data = yaml.safe_load(src)
            utils.update_metadata(data, name=args.service, namespace=args.namespace)
            print("Deploying {}".format(dst_file))
            yaml.dump(data, dst, default_flow_style=False)

    utils.apply(dst_file)

    src_file = os.path.join(os.getcwd(), "deploy/mariadb/mariadb-deployment.yaml")
    dst_file = os.path.join(os.getcwd(), "build/mariadb-deployment.yaml")
    with open(src_file, "r") as src:
        with open(dst_file, "w+") as dst:
            data = yaml.safe_load(src)
            utils.update_metadata(data, name=args.service, namespace=args.namespace)
            print("Deploying {}".format(dst_file))
            yaml.dump(data, dst, default_flow_style=False)
    utils.apply(dst_file)

    src_file = os.path.join(os.getcwd(), "deploy/mariadb/mariadb-storage.yaml")
    dst_file = os.path.join(os.getcwd(), "build/mariadb-storage.yaml")
    with open(src_file, "r") as src:
        with open(dst_file, "w+") as dst:
            data = src.read()
            try:
                size = utils.check_output(
                    f"kubectl -n {NAMESPACE} get persistentvolumeclaims mariadb-pv-claim " +
                    "-o=jsonpath='{.status.capacity.storage}'")
                print("Using existing disk size", size)
            except:
                size = "10Gi"
                print("Using default size", size)
            data = data.replace("REPLACE_STORAGE", size)
            data = yaml.safe_load(data)
            utils.update_metadata(data, name=args.service, namespace=args.namespace)
            print("Deploying {}".format(dst_file))
            yaml.dump(data, dst, default_flow_style=False)

    utils.apply(dst_file)


if __name__ == "__main__":
    main()
